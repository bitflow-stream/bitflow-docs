// GO111MODULE=auto go run generate.go
package main

import (
	"io/ioutil"
	"log"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

const (
	template_file = "nav-template.yml"
	md_suffix     = ".md"
)

var (
	url_regex    = regexp.MustCompile("^https://bitflow.readthedocs.io/(projects/([^/]+)/)?en/latest/(([^/]+)/)?$")
	main_project = "bitflow-docs"
	projects     = []string{"bitflow4j", "bitflow-antlr-grammars", "go-bitflow-collector", "go-bitflow", "python-bitflow", main_project}
)

type Nav struct {
	Nav []NavMenu
}

type NavMenu map[string][]NavItem
type NavItem map[string]string

func main() {
	var nav Nav
	loadYaml(template_file, &nav)

	for _, project := range projects {
		projectFile := project + ".yml"
		log.Println()
		log.Println("Handling project:", project)

		var mkdocs map[string]interface{}
		loadYaml(projectFile, &mkdocs)

		projectNav := cloneNav(&nav)
		patchNav(projectNav, project)

		mkdocs["nav"] = projectNav.Nav
		data, err := yaml.Marshal(mkdocs)
		if err != nil {
			log.Fatalln("Failed to marshall patched mkdocs file:", err)
		}

		err = ioutil.WriteFile(projectFile, data, 0664)
		err = nil
		if err != nil {
			log.Fatalln("Failed to write file", projectFile, ":", err)
		}
		log.Println("Written file:", projectFile)
	}
}

func loadYaml(filename string, target interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Failed to read YAML file", filename, ":", err)
	}
	err = yaml.Unmarshal(data, target)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Loaded YAML file", filename)
}

func cloneNav(nav *Nav) *Nav {
	cloneNavItem := func(item NavItem) NavItem {
		res := make(NavItem, len(item))
		for key, val := range item {
			res[key] = val
		}
		return res
	}
	cloneNavMenu := func(menu NavMenu) NavMenu {
		res := make(NavMenu, len(menu))
		for key, items := range menu {
			itemsCopy := make([]NavItem, len(items))
			for itemIndex, item := range items {
				itemsCopy[itemIndex] = cloneNavItem(item)
			}
			res[key] = itemsCopy
		}
		return res
	}

	res := &Nav{
		Nav: make([]NavMenu, len(nav.Nav)),
	}
	for i, menu := range nav.Nav {
		res.Nav[i] = cloneNavMenu(menu)
	}
	return res
}

func patchNav(nav *Nav, project string) {
	for menuIndex, menuItem := range nav.Nav {
		if len(menuItem) != 1 {
			log.Fatalf("Menu item %v for project %v has unexpected length %v", menuIndex, project, len(menuItem))
		}

		// Unpack single value
		var menuName string
		var menuEntries []NavItem
		for menuName, menuEntries = range menuItem {
		}

		log.Println("  Menu item:", menuName)
		for entryIndex, menuEntry := range menuEntries {
			if len(menuEntry) != 1 {
				log.Fatalf("    Entry %v has unexpected length %v", entryIndex, len(menuEntry))
			}

			// Unpack single value
			var entryName string
			var entry string
			for entryName, entry = range menuEntry {
			}

			match := url_regex.FindStringSubmatch(entry)
			if len(match) == 0 {
				log.Printf("    '%v': %v: External link, unchanged", entryName, entry)
			} else {
				projectName := match[2]
				if projectName == "" {
					projectName = main_project
				}
				pageName := match[4]
				if pageName == "" {
					pageName = "index"
				}

				if projectName == project {
					newEntry := pageName + md_suffix
					log.Printf("  ! '%v': Project-internal link, changing from: %v to: %v", entryName, entry, newEntry)
					menuEntry[entryName] = newEntry
				} else {
					log.Printf("    '%v': Link to project %v (page %v) %v", entryName, projectName, pageName, entry)
				}
			}
		}
	}
}

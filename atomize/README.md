Atomize
---
```
Usage: 
atomize [options] <yamlstream>
echo <yamlstream> | atomize [options]

Options:
  -h, --help            show this help message and exit
  -s, --nosort          Skip sorting logic
  -k, --nokustomize     Skip building kustomize files
  -c CUSTOMSORT, --customsort=CUSTOMSORT
                        File path for additional sorting data (yaml file)
  -n RESULTDIR, --name=RESULTDIR
                        Customize the name of the output folder
```

Inputs a multi-part yaml file consisting of multiple kubernetes resources, splits them into individual files with names based on the resource names, and optionally (by default) organizes them by category type and generates viable if simple kustomization.yaml files for easy deployment.

The script's output will be stored in a new directory within the current working directory. It can be manually named by the user by passing the -n or --name flag, but if this is not provided the name will match the name of the input file, if there is one, or will simply be named 'atomized'.

The purpose of the script is to provide a quick and easy way to make sense of the kubernetes code that is provided by outside sources such as helm or istioctl, and reconfigure the provided code into a consistent easily read format that complies with our linting standards and can be easily replicated.

Categorization is defined by the kubernetes resource kind in combination with data from the configuration file .atomize.yaml in the user's home directory. If this file does not exist it will be created, but the file can be edited to add additional resource types or for custom organization.

Simplified example config:
```
---
categories:
  ordered:
  - base
  - crd
  base:
    folder: .
    kinds:
    - Namespace
    - PriorityClass
  crd:
    folder: crd
    kinds:
    - CustomResourceDefinition
```

The configuration file must contain an entry for categories, and those categories must contain an 'ordered' array listing each of the categories by name. Those categories should each have an entry in the file that defines the folder they will be sorted into and an array of the different kubernetes resources that they should contain.

Any resources that are not defined will be sorted into an "unknown" category folder. If any unknown types are discovered or added (say from newly introduced operators with additional CDR entries) they should be added into the appropriate list.

Additional sorting can be achieved by providing a sorting config with the -c option.

Simplified example Config:
```
---
sorting:
  folders:
    - addons
    - addons/grafana
    - istiod
  labels:
    app-grafana: 'addons/grafana'
    app-istiod: 'istiod'
    app-galley: 'istiod'
  resources:
    istio-ConfigMap: 'istiod'
```

The sorting config must contain a 'sorting' entry that contains an array of folders, and additional mappings for labels and / or resources.

The folders list should explicitly contain any folders that you need to have created for your folder structure aside from those that would normally be generated according to the main config file. This includes any base folders (like 'addons' above).

Resources will still be further sorted into their usual categories within these new folders.

The labels list expects labels in the format of key-value (so kubernetes label app: grafana would match app-grafana above). The value given is the folder that you wish to sort those labels into. In the example above anything with the label app: grafana will end up under the 'addons/grafana' folder, while anything labeled app: isitod or app: gally will be placed in the istiod folder.

The resources list can be used to sort specific resources that do not have reliable or correct labeling when they are generated from their source. The key here should be the specific name of the resource, followed by a hyphen, and then the resource kind. This format matches the generated filename of the resources.

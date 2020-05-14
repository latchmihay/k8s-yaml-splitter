Atomize
---
```
Usage: 
atomize [options] <yamlstream>
echo <yamlstream> | atomize [options]

Options:
  -h, --help         show this help message and exit
  -s, --nosort       
  -k, --nokustomize 
```

Inputs a multi-part yaml file consisting of multiple kubernetes resources, splits them into individual files with names based on the resource names, and optionally (by default) organizes them by category type and generates viable if simple kustomization.yaml files for easy deployment.

Newly created folders and files will be placed in the current working directory.

The purpose of the script is to provide a quick and easy way to make sense of the kubernetes code that is provided by outside sources such as helm or istioctl, and reconfigure the provided code into a consistent easily read format that complies with our linting standards.

Categorization is defined by the "kinds" listed at the top of the script. Any resources that are not defined will be sorted into an "unknown" category folder. If any unknown types are discovered or added (say from newly introduced operators with additional CDR entries) they should be added into the appropriate list.
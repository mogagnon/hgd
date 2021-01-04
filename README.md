# Hosted Graphite Deploy
Hosted Graphite Deploy is a simple CLI that manages dashboards in hosted graphite. The main purpose of hgd is to automate the process of 
deploying dashboard in multiple instances of Hosted Graphite.

## Usage
````
Usage:
  hgd [command]

Available Commands:
  delete      Delete a Dashboard.
  export      Export a dashboard to Hosted Graphite from a json file. Will update if the dashboard already exist.
  help        Help about any command
  import      Import a dashboard locally.

Flags:
  -h, --help           help for hgd
  -t, --token string   The token for Hosted Graphite.

Use "hgd [command] --help" for more information about a command.

```` 
# Grizzly Lister: A tool for listing Grafana Dashboards and Folders

[Grizzly](https://github.com/grafana/grizzly) is a tool for declaratively managing
observability resources within Grafana Labs' systems.

This Grizzly Lister provides functionality that will eventually be rolled into
Grizzly itself. In this case, the ability to list resources found within a
Grafana instance. The purpose of this is to facilitate the migration of resources
between Grafana instances.

List IDs for folders on a Grafana instance:
```
grr-list folders
```
List dashboard UIDs within a specific folder:
```
grr-list dashboards <folder-id>
```

List names for all datasources on a Grafana instance:
```
grr-list datasources
```

See Grafana's [Authentication API
docs](https://grafana.com/docs/grafana/latest/http_api/auth/) for more info on
authentication.

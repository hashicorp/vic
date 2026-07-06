schema_version = 1

project {
  license        = "MPL-2.0"
  copyright_year = 2019

  header_ignore = [
    ".codecov.yml",
    ".drone.local.yml",
    ".drone.yml",
    "cmd/vic-machine-server/Dockerfile",
    "demos/compose/elk-app/docker-compose-elk.yml",
    "infra/integration-image/Dockerfile",
    "demos/compose/voting-app/docker-compose.yml",
    "demos/compose/webserving-app/docker-compose.yml",
    "isos/vicadmin/css/style.css",
    "isos/vicadmin/css/fontello.css",
    "isos/vicadmin/css/clarity-ui.min.css",
    "vendor/**",
    "tests/**",
  ]
}

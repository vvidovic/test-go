FROM golang:1.12

LABEL "com.github.actions.name"="go dpcmder build"
LABEL "com.github.actions.description"="Build Go DataPower Commander"
LABEL "com.github.actions.icon"="package"
LABEL "com.github.actions.color"="blue"

LABEL "maintainer"="Vedran Vidović <vvidovic@gmail.com>"

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

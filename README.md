# Gaia

## About Gaia
Gaia is a Remote Monitoring app for Linux devices. It allows a user to monitor several linux machines through a web browser.

## Technical Details
### Technical Overview
Gaia is split up into 3 parts:
1. A `Golang`-based "Sub Node" running on each of the linux machines being monitored. The Sub Node collects information about the machine it's running on (e.g. battery level, IP address, hostname, etc.).
2. A `Golang`-based "Master Node" running on a server that the user (i.e. you) own. The Master Node is responsible for connecting to all Sub Nodes and collecting their information.
3. A `Node.js` app that pulls data from the Master Node (via REST) and presents it in a browser as a table.

### Diagram
![alt text](https://raw.githubusercontent.com/itsamishra/Gaia/master/GaiaDiagram.png)

### Git Standards
Git branches are in the format {Person Committing}/{Project Card Number}/{Brief Description}. An example branch looks like:

`git branch ash/19131340/CreateBasicSubNode`

### Requirements
#### Master Node Requirements
`Golang 1.11+`
#### Sub Node Requirements
`Golang 1.11+`, `bash`, `upower`, `imagemagick`, `nmap`
#### Web Browser Requirements
`Node.js 11.11.0+`
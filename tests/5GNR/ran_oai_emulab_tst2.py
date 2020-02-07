#!/usr/bin/python

"""
5G-NR OAI RAN Scenario: gNodeB and UE with 2USRPX310
"""

import geni.portal as portal
import geni.rspec.pg as rspec

request = portal.context.makeRequestRSpec()

#Node0 @gNB
node0 = request.RawPC( "node0" )
node0.hardware_type = "x310"
node0.component_manager_id = "urn:publicid:IDN+emulab.net+authority+cm"

#Modify disk_image for node0:
node0.disk_image = "urn:publicid:IDN+emulab.net+image+PowderTeam:update-nr-image"
#Create temporary storage for node0:
bs0 = node0.Blockstore("bs0", "/mytempdata")
bs0.size = "10GB"


#Node1 @UE

node1 = request.RawPC( "node1" )
node1.disk_image = "urn:publicid:IDN+emulab.net+image+PowderTeam:update-nr-image"
node1.hardware_type = "x310"
node1.component_manager_id = "urn:publicid:IDN+emulab.net+authority+cm"
bs1 = node1.Blockstore("bs1", "/mytempdata")
bs1.size = "10GB"

#Add a link between the nodes:
wiredlink = request.Link("link")
wiredlink.addNode(node0)
wiredlink.addNode(node1)


portal.context.printRequestRSpec()

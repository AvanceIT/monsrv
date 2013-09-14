monsrv
======

Server end of Server monitoring system

Overview
--------

This monitor suite is designed to run clientside and report back to a remote central
server. The client runs monitor which wakes after a configured interval and runs
any checks required. If it finds a problem, it sends a message to the remote server
process monsrv. This message is an XML record which has been encrypted to protect
the data even if sent over the internet.

This is provided as Open Source Free software on an "as is" basis and licensed using
GPLv3. No liability is accpeted or implied.

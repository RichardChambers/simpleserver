# simpleserver
test server in go with JavaScript client

This repository is to experiment with having a simple server written in Go
that will provide services to point of sale.

The goal is to have a central server to which a point of sale can send
transaction data and then other systems or services can pull a copy of
the transaction data for some purpose.

Current functionality (May 1, 2016)

Working on a simple Kitchen Display System which has a client running in
a browser such as Chrome with JavaScript pulling the information to display
from the simpleserver.


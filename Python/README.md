# Python language

Code and docs related to the use of the Python programming language to
access the REMS REST API.

All of the code shown here will be using the REMS test instance provided
by the REMS developers:

- the instance is at
[https://rems-demo.rahtiapp.fi](https://rems-demo.rahtiapp.fi)
- the API swagger page is at
[https://rems-demo.rahtiapp.fi/swagger-ui/index.html](https://rems-demo.rahtiapp.fi/swagger-ui/index.html)
- login accounts are at [https://rems-demo.rahtiapp.fi/fake-login](https://rems-demo.rahtiapp.fi/fake-login)

## example1

This is a toy command line application that makes a request to a single
REMS endpoint (`my-applications`) using the REST API.
Everything is hard coded so it is not a useful starting point for 
production code but it does show the basics of using Python and of
using the api-key based method of authentication.
It relies only on the standard requests module so there are no 
external dependencies.

Usage is just:

python example1.py
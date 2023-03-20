# Example 1 - basic go application

This example shows how to use go language to contact a REMS REST API.
It has the api-key and user-id hard-coded in which is a very bad idea
so please DO NOT do this in production.

To compile this program you will need go version 1.20 or later. To see
if you have go installed and what version, try:

```
go version
```

You will hopefully see something like:

```
go version go1.20.1 darwin/arm64
```

Now change directory into wherever you cloned this source code to
and type the following command:

```
go build
```

You should now find a `ex1` executable sitting in the directory. It
contacts the REMS demo instance provided by the REMS creators and it has
the authentication information hard coded in (yes I already said I know
that is a bad idea) so all you need to do is invoke it:

```
./ex1
```

The only thing written to STDOUT is the JSON response body so you can
safely pipe this through jq to get a more intelligible pretty-printed
 view of the JSON retruned by the REMS server:

```
./ex1 | jq
```

`ex1` does some very basic logging so you might want to redirect STDERR to
a file. If yo don't redirect the logging, it will appear on the terminal
but it will not affect `jq` because the logging goes to STDERR and it's
the output of STDOUT that gets piped to jq.

```
./ex1 2> my.log | jq
```

The output is a list of applications that the user can see and it's very
long but the start should look something like:

```[
  {
    "application/workflow": {
      "workflow/id": 1,
      "workflow/type": "workflow/default"
    },
    "application/external-id": "2023/8",
    "application/first-submitted": "2023-03-18T00:02:19.630Z",
    "application/blacklist": [],
    "application/id": 8,
    "application/duo": {
      "duo/matches": []
    },
    "application/applicant": {
      "userid": "RDapplicant1@funet.fi",
      "name": "RDapplicant1 REMSDEMO1",
      "email": "RDapplicant1.test@test_example.org",
      "organizations": [
        {
          "organization/id": "default"
        }
      ]
    },
    "application/todo": "waiting-for-review",
    "application/members": [],
    "application/resources": [
      {
        "catalogue-item/end": null,
        "catalogue-item/expired": false,
        "catalogue-item/enabled": true,
        "resource/id": 1,
        "catalogue-item/title": {
          "en": "Default workflow",
          "fi": "Oletustyövuo",
          "sv": "Standard arbetsflöde"
        },
        "catalogue-item/infourl": {
          "en": "http://www.google.com",
          "fi": "http://www.google.fi",
          "sv": "http://www.google.se"
        },
        "resource/ext-id": "urn:nbn:fi:lb-201403262",
        "catalogue-item/start": "2023-03-18T00:02:18.054Z",
        "catalogue-item/archived": false,
        "catalogue-item/id": 3
      }
    ],
    ...
```

`...` is where I've elided the rest of the output for space
cponsiderations.

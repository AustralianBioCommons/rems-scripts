# Using curl to access the REMS REST API

Examples of use of curl commands to access the REMS REST API.

In this scenario we are making an application for a REMS catalog item. The item uses a REMS form with a field to which a user must add an attachment. We can use the API to upload a file, attach it to the application, and then set the attachment as the value of the field. 

We have an API key for authentication, and know the applicant's user id, the application id, the form id and the field id, all of which we save in environment variables.

We proceed in two steps. Note that the REMS test instance's API swagger page can be used to explore the API services which will be used:

- [https://rems-demo.rahtiapp.fi/swagger-ui/index.html#/applications/post_api_applications_add_attachment](https://rems-demo.rahtiapp.fi/swagger-ui/index.html#/applications/post_api_applications_add_attachment)
- [https://rems-demo.rahtiapp.fi/swagger-ui/index.html#/applications/post_api_applications_save_draft](https://rems-demo.rahtiapp.fi/swagger-ui/index.html#/applications/post_api_applications_save_draft)


## 1. upload the attachment

```
curl -X 'POST' \
  "${SERVER}/api/applications/add-attachment?application-id=${APPLICATION_ID}" \
  -H "accept: application/json" \
  -H "x-rems-api-key: $API_KEY" \
  -H "x-rems-user-id: $USER_ID" \
  -H "Content-Type: multipart/form-data" \
  -F "file=@${PATH_TO_FILE_TO_UPLOAD}"

```
If successful this returns a json string which includes the id of the attachment.

## 2. associate the attachment with the field in the form
```
curl -X 'POST' \
  "${SERVER}/api/applications/save-draft" \
  -H "accept: application/json" \
  -H "x-rems-api-key: $API_KEY" \
  -H "x-rems-user-id: $USER_ID" \
  -H "Content-Type: application/json" \
  -d '{"application-id": '${APPLICATION_ID}', "field-values": [{ "form": '${FORM_ID}', "field": "'${FIELD_ID}'", "value": "'${ATTACHMENT_ID}'" }]}'
```

See [https://github.com/CSCfi/rems/discussions/3252#discussioncomment-8142853](this discussion).

# txt2gdoc

Creates Google Document from text input

**Note**: to use this project you will need to provide `credentials.json` and `token.json` files to allow `txt2gdoc` to
create Google Documents in your account. The access token need to contain `/auth/documents` scope.

If you do not have credentials nor token from Google. You could create and configure your own Google Cloud OAuth client.
You can follow the instructions below on how to enable OAuth authorization.

If you have credentials file downloaded (as described in *How to enable OAuth authorization*) you can create a token file
by using following command:

`go run github.com/sebast26/txt2gdoc/cmd/txt2gdoc-setup`

this will create `token.json` file.

## How to enable OAuth authorization

1. Go to the [Google Cloud Platform Console](https://console.cloud.google.com/) and from the projects list, select a project or create a new one.
2. Disable all APIs and enable `Google Docs API`.
3. Follow [this guide](https://support.google.com/cloud/answer/6158849) from Google to create OAuth client ID. As a scope for your application
you may select `/auth/documents` from Google Docs API. Add you email to the tests user group. For an application type you have to select `Computer`.
4. Download and save your client credentials as `credentials.json` file.

## How to run

`txt2gdoc` reads stdin and creates Google Document with the content given. To execute the program please run:

`cat README.md | go run github.com/sebast26/txt2gdoc/cmd/txt2gdoc`
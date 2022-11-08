# txt2gdoc

Creates Google Document from text input

**Note**: to use this project you need to configure Google Cloud OAuth client ID. This is required for `txt2gdoc` to create Google Docs in your account.
You can follow the instructions below on how to enable OAuth authorization.

## How to enable OAuth authorization

1. Go to the [Google Cloud Platform Console](https://console.cloud.google.com/) and from the projects list, select a project or create a new one.
2. Disable all APIs and enable `Google Docs API`.
3. Follow [this guide](https://support.google.com/cloud/answer/6158849) from Google to create OAuth client ID. As a scope for your application
you may select `/auth/documents` from Google Docs API. Add you email to the tests user group. For an application type you have to select `Computer`.
4. Download and save your client credentials as `credentials.json` file.


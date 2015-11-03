[![Build Status](https://travis-ci.org/albertoleal/features.png?branch=master)](https://travis-ci.org/albertoleal/features)
[![Coverage Status](https://coveralls.io/repos/albertoleal/features/badge.svg?branch=master&service=github)](https://coveralls.io/github/albertoleal/features?branch=master)

# Features

Features helps you to determine whether or not a feature should be made available to one or more users.
You can choose one of the following ways:

  * Enable a feature only for a specific set of users;
  * Enable or disabled for a percentage of users;
  * Disable a feature.

Also known as: feature switch, feature flag, feature toggle, ...

## Storage
  * Features is completely storage agnostic. You should be able to use your own storage, you just need to implement the [`Engine`](https://github.com/albertoleal/features/blob/master/engine/engine.go) interface.
  * This library comes with an in memory store, but it's basically used for testing. You should not use this in production.

## How it works

  Create an instance of Features type passing the storage as an argument:

  ```golang
  Features := features.New(memory.New())
  ```

  Create an instance of Feature type with the desired parameters and save it:

  ```golang
  feature := engine.FeatureFlag{
    Key:     "Feature X",
    Enabled: false,
    Users:   []*engine.User{&engine.User{Id: "alice@example.org"}},
    Percentage: 10,
  }
  Features.Save(feature)
  ```

  Check if a feature is **enabled**:

  ```golang
  Features.IsEnabled("Feature X")
  ```

  Check if a feature is **disabled**:

  ```golang
  Features.IsDisabled("Feature X")
  ```

  Execute an anonymous function if the feature is **enabled**:
  ```golang
  Features.With("Feature X", func() {
    fmt.Println("`Feature X` is enabled!")
  })
  ```

  Execute an anonymous function if the feature is **disabled**:
  ```golang
  Features.Without("Feature X", func() {
    fmt.Println("`Feature X` is disabled!")
  })
  ```

  Check if a feature is active for a particular user:

  ```golang
  Features.UserHasAccess("Feature X", "alice@example.org")
  ```

## User percentages

If you're rolling out a feature, you might want to enable it for a percentage of your users. There are two ways you can achieve that: Either enabling the feature for a percentage of users or disabling that.

  * **Enable** the feature for a percentage of users:
  ```golang
  feature := engine.FeatureFlag{
    Key:     "Feature X",
    Enabled: true,
    Percentage: 10,
  }
  ```

  * **Disable** the feature for a percentage of users:
  ```golang
  feature := engine.FeatureFlag{
    Key:     "Feature X",
    Enabled: false,
    Percentage: 10,
  }
  ```

## Specific users

You might want to enable a feature for a set of users to try out and give feedback on before it's rolled out for everyone. To achieve this, you need to specify the users by doing the following:

  ```golang
  feature := engine.FeatureFlag{
    Key:     "Feature X",
    Enabled: true,
    Users:   []*engine.User{&engine.User{Id: "alice@example.org"}},
  }
  ```

## HTTP API
  Features comes with an HTTP API that's used by the SDKs, but you **should not** expose that to the world because there's no authentication layer.

### API endpoints
  - [`POST` /features](#create-a-new-feature-flag) - Create a new feature flag
  - [`PUT` /features/:feature-key](#update-an-existing-feature-flag) - Update an existing feature flag
  - [`DELETE` /features/:feature-key](#delete-an-existing-feature-flag) - Delete an existing feature flag
  - [`GET` /features](#list-all-feature-flags) - List all feature flags
  - [`GET` /features/:feature-key](#get-an-existing-feature-flag) - Get an existing feature flag

  #### Create a new feature flag
  ```bash
  curl -i http://localhost:8000/features -XPOST -d '{"key": "login-via-email", "percentage": 90, "enabled": true, "users": [{"id": "alice@example.org"}]}'
  ```
  Response:
  ```bash
  HTTP/1.1 201 Created
  Content-Type: application/json
  Date: Tue, 03 Nov 2015 23:21:07 GMT
  Content-Length: 94

  {"enabled":true,"key":"login-via-email","users":[{"id":"alice@example.org"}],"percentage":90}
  ```

  #### Update an existing feature flag
  ```bash
  curl -i http://localhost:8000/features/login-via-email -XPUT -d '{"key": "login-via-email", "percentage": 90, "enabled": false, "users": [{"id": "bob@example.org"}]}'
  ```
  Response:
  ```bash
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Tue, 03 Nov 2015 23:22:05 GMT
  Content-Length: 93

  {"enabled":false,"key":"login-via-email","users":[{"id":"bob@example.org"}],"percentage":90}
  ```

  #### Delete an existing feature flag
  ```bash
  curl -i http://localhost:8000/features/login-via-email -XDELETE
  ```
  Response:
  ```bash
  HTTP/1.1 204 No Content
  Content-Type: application/json
  Date: Tue, 03 Nov 2015 23:24:46 GMT
  ```

  #### Get an existing feature flag
  ```bash
  curl -i http://localhost:8000/features/login-via-email
  ```
  Response:
  ```bash
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Tue, 03 Nov 2015 23:26:22 GMT
  Content-Length: 94

  {"enabled":true,"key":"login-via-email","users":[{"id":"alice@example.org"}],"percentage":90}
  ```

  #### List all feature flags
  ```bash
  curl -i http://localhost:8000/features
  ```
  Response:
  ```javascript
  {
    "items": [
      {
        "percentage": 90,
        "enabled": true,
        "users": [
          {
            "id": "alice@example.org"
          }
        ],
        "key": "login-via-email"
      }
    ],
    "item_count": 1
  }
  ```

[![Build Status](https://travis-ci.org/albertoleal/features.png?branch=master)](https://travis-ci.org/albertoleal/features)

# Features


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

  Check if a feature is **active**:

  ```golang
  Features.IsActive("Feature X")
  ```

  Check if a feature is **inactive**:

  ```golang
  Features.IsInactive("Feature X")
  ```

  Execute an anonymous function if the feature is **active**:
  ```golang
  Features.With("Feature X", func() {
    fmt.Println("`Feature X` is enabled!")
  })
  ```

  Execute an anonymous function if the feature is **inactive**:
  ```golang
  Features.Without("Feature X", func() {
    fmt.Println("`Feature X` is disabled!")
  })
  ```

  Check if a feature is active for a particular user:

  ```golang
  Features.UserHasAccess("Feature X", "alice@example.org")
  ```

## User Percentages

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

## Specific Users

You might want to enable a feature for your team to try out and give feedback on before it's rolled out for everyone. To achieve this, you need to specify the users by doing the following:

  ```golang
  feature := engine.FeatureFlag{
    Key:     "Feature X",
    Enabled: true,
    Users:   []*engine.User{&engine.User{Id: "alice@example.org"}},
  }
  ```

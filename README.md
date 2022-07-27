# ocm-support-cli

`ocm-support-cli` is a tool that extends the `ocm-cli` by adding commands that are useful for engineers that deal support requests.

Prerequisites: 

* Have `ocm` command installed and have run `ocm login` command.
* Have advanced permissions provided by roles suhc as `UHCSupport`,`SDBSignalMonitor`, ...

## Install

### Option 1: Build from source
First clone the repository somewhere in your $PATH. A common place would be within your $GOPATH.

Example:

```
mkdir $GOPATH/src/github.com/openshift-online
cd $GOPATH/src/github.com/openshift-online
git clone git@github.com:openshift-online/ocm-support-cli.git
```

Next, cd into the ocm-support-cli folder and run `make install`. This command will build the `ocm-support` binary and place it in $GOPATH. As the binary has prefix `ocm-`, it becomes a plugin of `ocm`, and can be invoked by `ocm support`.

### Validating the Installation

You can use the following command to confirm that the tool was installed correctly:

`ocm support version`

## Usage

To see all available commands, run `ocm support -h`.

### Accounts

The `accounts` command gets information about or execute actions on accounts.

#### Finding an account

Use the `find` subcommand to find one or more accounts, passing as argument one of the following:

* accountID
* username
* email
* organizationID
* organizationExternalID
* organizationEBSAccountID

The following flags are available for `accounts find`:

```
--all                        If true, returns all accounts that matched the search instead of the first one only (default behaviour).
--fetchRegistryCredentials   If true, includes the account registry credentials.
--fetchRoles                 If true, includes the account roles.
--fetchLabels                If true, includes the account labels.
--fetchCapabilities          If true, includes the account capabilities.
-h, --help                   help for find
```

#### Examples

* Find an account by email `ocm support accounts find user@redhat.com`
* Find an account and include its roles in the results `ocm support accounts find [accountID] --fetchRoles`
* Find an account and include its labels in the results `ocm support accounts find [accountID] --fetchLabels`
* Find an account and include its capabilities in the results `ocm support accounts find [accountID] --fetchCapabilities`
* Find all accounts from an organization `ocm support accounts find [organizationID] --all`

### Organizations

The `organizations` command gets information about or execute actions on organizations.

#### Finding an organization

Use the `find` subcommand to find one or more organizations, passing as argument one of the following:

* organizationID
* organizationExternalID
* organizationEBSAccountID

The following flags are available for `organizations find`:

```
--all                  If true, returns all organizations that matched the search instead of the first one only (default behaviour).
--fetchQuota           If true, includes the organization quota.
--fetchSubscriptions   If true, includes the organization subscriptions.
--fetchLabels          If true, includes the organization labels.
--fetchCapabilities    If true, includes the organization capabilities.
-h, --help             help for find
```

#### Examples

* Find an organization by its externalID: `ocm support organizations find [organizationExternalID]`
* Find an organization and include its subscriptions: `ocm support organizations find [organizationID] --fetchSubscriptions`
* Find an organization and include its labels `ocm support organizations find [organizationID] --fetchLabels --fetchCapabilities`
* Find an organization and include its capabilities `ocm support organizations find [organizationID] --fetchCapabilities`
* Find an organization and include its quota: `ocm support organizations find [organizationID] --fetchQuota`

### RegistryCredentials

The `registryCredentials` command gets information about or execute actions on registry credentials.

#### Creating registry credentials

Use the `create` subcommand to create registry credentials for current account. 

#### Displaying registry credentials

Use the `show` subcommand to to see registry credentials, passing accountID.

#### Deleting registry credentials

Use the `delete` subcommand to to delete a specific registry credential, or all registry credentials, for the passed accountID.

The following flags are available for `registryCredentials delete`:

```
--all                        If true, deletes all registry credentials for the given account ID.
-h, --help                   help for delete
```

#### Examples

* Create registry credentials `ocm support registryCredentials create`
* Show registry credentials for a specific account `ocm support registryCredentials show [accountID]`
* Delete a specific registry credential for a specific account `ocm support registryCredentials delete [accountID] [registryCredentialID]`
* Delete all registry credentials for a specific account `ocm support registryCredentials delete [accountID] --all`

### Create

The `create` command creates the given resource with provided key and value.

Available values for creating resources are `accountLabel|organizationLabel|subscriptionLabel|accountCapability|organizationCapability|subscriptionCapability`

Pass key and value for creating a label, and pass a valid key for creating a capability

The following flags are available for `create`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

#### Examples

* Add label to an account `ocm support create accountLabel [accountID] [key] [value]`
* Add label to an organization `ocm support create organizationLabel [orgID] [key] [value]`
* Add label to a subscription with internal flag as false `ocm support create subscriptionLabel [subscriptionID] [key] [value] --external`
* Add capability to an account `ocm support create accountCapability [accountID] [capabilityKey]`
* Add capability to an organization `ocm support create organizationCapability [orgID] [capabilityKey]`
* Add capability to a subscription `ocm support create subscriptionCapability [subscriptionID] [capabilityKey]`
# ocm-support-cli

`ocm-support-cli` is a tool that extends the `ocm-cli` by adding commands that are useful for engineers that deal support requests.

Prerequisites: 

* Have `ocm` command installed and have run `ocm login` command.
* Have advanced permissions provided by roles suhc as `UHCSupport`,`SDBSignalMonitor`, ...

## Install

### Option 1: Download binary

Download the latest binary file from the [release page](https://github.com/openshift-online/ocm-support-cli/releases).

For Linux, download `ocm-support_linux_amd64`, rename it to `ocm-support` and put it to $PATH. For example:
~~~
$ sudo cp ocm-support_linux_amd64 /usr/bin/ocm-support
$ sudo chmod 0755 /usr/bin/ocm-support
~~~

For MacOS, download `ocm-support_darwin_amd64`, rename it to `ocm-support` and put it to $PATH. For example:
~~~
$ sudo cp ocm-support_darwin_amd64 /usr/local/bin/ocm-support
$ sudo chmod 0755 /usr/local/bin/ocm-support
~~~

### Option 2: Build from source
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

### Get

The `get` command gets the given resource.

#### Getting an account

Use the `accounts` subcommand to get one or more accounts, passing as argument one of the following:

* accountID
* username
* email
* organizationID
* organizationExternalID
* organizationEBSAccountID

Pass the search criteria as an optional second argument.

The following flags are available for `get accounts`:

```
--first                      If true, returns only the first accounts that matched the search instead of all of them (default behaviour).
--fetchRegistryCredentials   If true, includes the account registry credentials.
--fetchRoles                 If true, includes the account roles.
--fetchLabels                If true, includes the account labels.
--fetchCapabilities          If true, includes the account capabilities.
-h, --help                   help for get
```

##### Examples

* Get the first account by email `ocm support get accounts user@redhat.com --first`
* Get the account and include its roles in the results `ocm support get accounts [accountID] --fetchRoles`
* Get the account and include its registry credentials `ocm support get accounts [username] --fetchRegistryCredentials`
* Get all accounts for an organizationExternalID and include its labels in the results `ocm support get accounts [organizationExternalID] --fetchLabels`
* Get all accounts for an organizationEBSAccountID and include its capabilities in the results `ocm support get accounts [organizationEBSAccountID] --fetchCapabilities`
* Get all accounts from an organization `ocm support get accounts [organizationID]`

#### Getting an organization

Use the `organizations` subcommand to get one or more organizations, passing as argument one of the following:

* organizationID
* organizationExternalID
* organizationEBSAccountID

Pass the search criteria as an optional second argument.

The following flags are available for `get organizations`:

```
--first                If true, returns only the first accounts that matched the search instead of all of them (default behaviour).
--fetchQuota           If true, includes the organization quota.
--fetchSubscriptions   If true, includes the organization subscriptions.
--fetchLabels          If true, includes the organization labels.
--fetchCapabilities    If true, includes the organization capabilities.
--fetchSkus            If true, returns all the resource quota objects for the organization.
-h, --help             help for get
```

##### Examples
* Get the first organization by its externalID: `ocm support get organizations [organizationExternalID] --first`
* Get the organization and include its subscriptions: `ocm support get organizations [organizationID] --fetchSubscriptions`
* Get all organizations for an organizationExternalID and include its labels `ocm support get organizations [organizationExternalID] --fetchLabels`
* Get all organizations for an organizationEBSAccountID and include its capabilities `ocm support get organizations [organizationEBSAccountID] --fetchCapabilities`
* Get the first organization and include its quota: `ocm support get organizations [organizationID] --first --fetchQuota`
* Get all organizations for an organizationExternalID and include its SKUs: `ocm support get organizations [organizationExternalID] --fetchSkus`

#### Getting a subscription

Use the `subscriptions` subcommand to get one or more subscriptions, passing as argument one of the following:

* subscriptionID
* clusterID
* externalClusterID
* organizationID

Pass the search criteria as an optional second argument.

The following flags are available for `get subscriptions`:

```
--first                     If true, returns only the first subscription that matches the search instead of all of them (default behaviour).
--fetchLabels               If true, includes the organization labels.
--fetchCapabilities         If true, includes the organization capabilities.
--fetchReservedResources    If true, returns a list of reserved resources for the subscriptions.
-h, --help                  help for get
```

##### Examples

* Get subscription by its ID: `ocm support get subscriptions [subscriptionID]`
* Get all subscriptions by ClusterID and include its labels `ocm support get subscriptions [clusterID] --fetchLabels`
* Get all subscriptions by ClusterID and include its capabilities `ocm support get subscriptions [clusterID] --fetchCapabilities`
* Get first subscription by its externalClusterID: `ocm support get subscriptions [externalClusterID] --first`
* Get all subscriptions by OrganizationID and include subscriptions that have Status as 'Reserverd' `ocm support get subscriptions [organizationID] "Status='Reserved'"`
* Get subscription by its ID and include its reserved resources: `ocm support get subscriptions [subscriptionID] --fetchReservedResources`

#### Getting registry credentials

Use the `registryCredentials` subcommand to to get registry credentials, passing accountID.

##### Examples

* Show registry credentials for a specific account `ocm support get registryCredentials [accountID]`

### Create

The `create` command creates the given resource.

#### Creating labels

##### Creating account labels

Use the `accountLabel` subcommand to assign a label to an account, passing a key and a value. 

The following flags are available for `create accountLabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

##### Creating organization labels

Use the `organizationLabel` subcommand to assign a label to an organization, passing a key and a value. 

The following flags are available for `create organizationLabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

##### Creating subscription labels

Use the `subscriptionLabel` subcommand to assign a label to a subscription, passing a key and a value. 

The following flags are available for `create subscriptionLabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

##### Examples

* Create a label to an account `ocm support create accountLabel [accountID] [key] [value]`
* Create a label to an organization `ocm support create organizationLabel [orgID] [key] [value]`
* Create a label to a subscription (with internal flag as false) `ocm support create subscriptionLabel [subscriptionID] [key] [value] --external`

#### Creating capabilities

##### Creating account capabilities

Use the `accountCapability` subcommand to assign a capability to an account, passing a valid capability key. 

##### Creating organization capabilities

Use the `organizationCapability` subcommand to assign a capability to an organization, passing a valid capability key. 

##### Creating subscription capabilities

Use the `subscriptionCapability` subcommand to assign a capability to a subscription, passing a valid capability key. 

##### Examples

* Create a capability to an account `ocm support create accountCapability [accountID] [capabilityKey]`
* Create a capability to an organization `ocm support create organizationCapability [orgID] [capabilityKey]`
* Create a capability to a subscription `ocm support create subscriptionCapability [subscriptionID] [capabilityKey]`

#### Creating registry credentials

Use the `registryCredentials` subcommand to create registry credentials for current account. 

##### Examples

* Create registryCredentials `ocm support create registryCredentials`

#### Creating role bindings

##### Creating application role bindings

Use the `applicationRoleBinding` subcommand to assign a role binding to an account at application level, passing a valid role id. 

##### Creating organization role bindings

Use the `organizationRoleBinding` subcommand to assign a role binding to an account at organization level, passing a valid role id.

##### Creating subscription role bindings

Use the `subscriptionRoleBinding` subcommand to assign a role binding to an account at subscription level, passing a valid role id.

##### Examples

* Create a role binding to an application `ocm support create applicationRoleBinding [accountID] [roleID]`
* Create a role binding to an organization `ocm support create organizationRoleBinding [accountID] [orgID] [roleID]`
* Create a role binding to a subscription `ocm support create subscriptionRoleBinding [accountID] [subscriptionID] [roleID]`

### Delete

The `delete` command deletes the given resource.

#### Deleting labels

##### Deleting account labels

Use the `accountLabel` subcommand to delete a label from an account, passing the label key. 

##### Deleting organization labels

Use the `organizationLabel` subcommand to delete a label from an organization, passing the label key. 

##### Deleting subscription labels

Use the `subscriptionLabel` subcommand to delete a label from a subscription, passing the label key. 

##### Examples

* Delete a label from an account `ocm support delete accountLabel [accountID] [key]`
* Delete a label from an organization `ocm support delete organizationLabel [orgID] [key]`
* Delete a label from a subscription `ocm support delete subscriptionLabel [subscriptionID] [key]`

#### Deleting capabilities

##### Deleting account capabilities

Use the `accountCapability` subcommand to delete a capability from an account, passing the valid capability key. 

##### Deleting organization capabilities

Use the `organizationCapability` subcommand to delete a capability from an organization, passing the valid capability key. 

##### Deleting subscription capabilities

Use the `subscriptionCapability` subcommand to delete a capability from a subscription, passing a valid capability key. 

##### Examples

* Delete a capability from an account `ocm support delete accountCapability [accountID] [capabilityKey]`
* Delete a capability from an organization `ocm support delete organizationCapability [orgID] [capabilityKey]`
* Delete a capability from a subscription `ocm support delete subscriptionCapability [subscriptionID] [capabilityKey]`

##### Deleting type independent capabilities

Use the `capability` subcommand to delete a capability by passing the ID, or provide filter value to search matching capabilities and delete them. By default the dry run flag will be enabled. Set `dryRun` flag to false to actually remove the resource.

The following flags are available for `delete capability`:

```
--filter                     If non-empty, filters and deletes the matching capabilities.
--dryRun                     If false, it will execute the delete command call in instead of a dry run.
-h, --help                   help for create
```

##### Examples

* Delete a capability by its ID `ocm support delete capability [capabilityID]`
* Delete all capabilities where key is 'capability.account.create_moa_clusters' (with no dry run) `ocm support delete capability --filter "key = 'capability.account.create_moa_clusters'" --dry-run=false`

#### Deleting registry credentials

Use the `registryCredentials` subcommand to to delete a specific registry credential, or all registry credentials, for the passed accountID.

The following flags are available for `registryCredentials delete`:

```
--all                        If true, deletes all registry credentials for the given account ID.
-h, --help                   help for delete
```

#### Examples

* Delete a specific registry credential for a specific account `ocm support delete registryCredentials [accountID] [registryCredentialID]`
* Delete all registry credentials for a specific account `ocm support delete registryCredentials [accountID] --all`

#### Deleting role bindings

##### Deleting application role bindings

Use the `applicationRoleBinding` subcommand to remove a role binding from an account at application level, passing a valid role id. 

##### Creating organization role bindings

Use the `organizationRoleBinding` subcommand to remove a role binding from an account at organization level, passing a valid role id.

##### Creating subscription role bindings

Use the `subscriptionRoleBinding` subcommand to remove a role binding from an account at subscription level, passing a valid role id.

##### Examples

* For the given account, delete a role binding created at application level using the roleID `ocm support delete applicationRoleBinding [accountID] [roleID]`
* For the given account, delete a role binding created at organization level using the roleID `ocm support delete organizationRoleBinding [accountID] [orgID] [roleID]`
* For the given account, delete a role binding created at subscription level using the roleID `ocm support delete subscriptionRoleBinding [accountID] [subscriptionID] [roleID]`

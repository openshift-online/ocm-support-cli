# ocm-support-cli

`ocm-support-cli` is a tool that extends the `ocm-cli` by adding commands that are useful for engineers that deal support requests.

Prerequisites: 

* Have `ocm` command installed and have run `ocm login` command.
* Have advanced permissions provided by roles suhc as `UHCSupport`,`SDBSignalMonitor`, ...

## Install

### Option 1: Download binary

Download the latest binary file from the [release page](https://github.com/openshift-online/ocm-support-cli/releases).

For Linux, download `ocm-support-linux-amd64`, rename it to `ocm-support` and put it to $PATH. For example:
~~~
$ sudo cp ocm-support-linux-amd64 /usr/bin/ocm-support
$ sudo chmod 0755 /usr/bin/ocm-support
~~~

For MacOS, download `ocm-support-darwin-amd64`, rename it to `ocm-support` and put it to $PATH. For example:
~~~
$ sudo cp ocm-support-darwin-amd64 /usr/local/bin/ocm-support
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
--fetch-registry-credentials If true, includes the account registry credentials.
--fetch-roles                If true, includes the account roles.
--fetch-labels               If true, includes the account labels.
--fetch-capabilities         If true, includes the account capabilities.
--fetch-export-control       If true, includes export control information.
-h, --help                   help for get
```

##### Examples

* Get the first account by email `ocm support get accounts user@redhat.com --first`
* Get the account and include its roles in the results `ocm support get accounts [accountID] --fetch-roles`
* Get the account and include its registry credentials `ocm support get accounts [username] --fetch-registry-credentials`
* Get the account and include export control information `ocm support get accounts [username] --fetch-export-control`
* Get all accounts for an organizationExternalID and include its labels in the results `ocm support get accounts [organizationExternalID] --fetch-labels`
* Get all accounts for an organizationEBSAccountID and include its capabilities in the results `ocm support get accounts [organizationEBSAccountID] --fetch-capabilities`
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
--fetch-quota          If true, includes the organization quota.
--fetch-subscriptions  If true, includes the organization subscriptions.
--fetch-labels         If true, includes the organization labels.
--fetch-capabilities   If true, includes the organization capabilities.
--fetch-skus           If true, returns all the resource quota objects for the organization.
-h, --help             help for get
```

##### Examples
* Get the first organization by its externalID: `ocm support get organizations [organizationExternalID] --first`
* Get the organization and include its subscriptions: `ocm support get organizations [organizationID] --fetch-subscriptions`
* Get all organizations for an organizationExternalID and include its labels `ocm support get organizations [organizationExternalID] --fetch-labels`
* Get all organizations for an organizationEBSAccountID and include its capabilities `ocm support get organizations [organizationEBSAccountID] --fetch-capabilities`
* Get the first organization and include its quota: `ocm support get organizations [organizationID] --first --fetch-quota`
* Get all organizations for an organizationExternalID and include its SKUs: `ocm support get organizations [organizationExternalID] --fetch-skus`

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
--fetch-labels               If true, includes the organization labels.
--fetch-capabilities         If true, includes the organization capabilities.
--fetch-reserved-resources    If true, returns a list of reserved resources for the subscriptions.
--fetch-roles                If true, returns the subscription roles.
-h, --help                  help for get
```

##### Examples

* Get subscription by its ID: `ocm support get subscriptions [subscriptionID]`
* Get all subscriptions by ClusterID and include its labels `ocm support get subscriptions [clusterID] --fetch-labels`
* Get all subscriptions by ClusterID and include its capabilities `ocm support get subscriptions [clusterID] --fetch-capabilities`
* Get first subscription by its externalClusterID: `ocm support get subscriptions [externalClusterID] --first`
* Get all subscriptions by OrganizationID and include subscriptions that have Status as 'Reserverd' `ocm support get subscriptions [organizationID] "Status='Reserved'"`
* Get subscription by its ID and include its reserved resources: `ocm support get subscriptions [subscriptionID] --fetch-reserved-resources`
* Get first subscription by its cluster ID and include its roles: `ocm support get subscriptions [clusterID] --first --fetch-roles`

#### Getting registry credentials

Use the `registrycredentials` subcommand to to get registry credentials, passing accountID.

##### Examples

* Show registry credentials for a specific account `ocm support get registrycredentials [accountID]`

### Create

The `create` command creates the given resource.

#### Creating labels

##### Creating account labels

Use the `accountlabel` subcommand to assign a label to an account, passing a key and a value. 

The following flags are available for `create accountlabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

##### Creating organization labels

Use the `organizationlabel` subcommand to assign a label to an organization, passing a key and a value. 

The following flags are available for `create organizationlabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

##### Creating subscription labels

Use the `subscriptionlabel` subcommand to assign a label to a subscription, passing a key and a value. 

The following flags are available for `create subscriptionlabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

##### Examples

* Create a label to an account `ocm support create accountlabel [accountID] [key] [value]`
* Create a label to an organization `ocm support create organizationlabel [orgID] [key] [value]`
* Create a label to a subscription (with internal flag as false) `ocm support create subscriptionlabel [subscriptionID] [key] [value] --external`

#### Creating capabilities

##### Creating account capabilities

Use the `accountcapability` subcommand to assign a capability to an account, passing a valid capability key. 

##### Creating organization capabilities

Use the `organizationcapability` subcommand to assign a capability to an organization, passing a valid capability key. 

##### Creating subscription capabilities

Use the `subscriptioncapability` subcommand to assign a capability to a subscription, passing a valid capability key. 

##### Examples

* Create a capability to an account `ocm support create accountcapability [accountID] [capabilityKey]`
* Create a capability to an organization `ocm support create organizationcapability [orgID] [capabilityKey]`
* Create a capability to a subscription `ocm support create subscriptioncapability [subscriptionID] [capabilityKey]`

#### Creating registry credentials

Use the `registrycredentials` subcommand to create registry credentials for current account. 

##### Examples

* Create registryCredentials `ocm support create registrycredentials`

#### Creating role bindings

##### Creating application role bindings

Use the `applicationrolebinding` subcommand to assign a role binding to an account at application level, passing a valid role id. 

##### Creating organization role bindings

Use the `organizationrolebinding` subcommand to assign a role binding to an account at organization level, passing a valid role id.

##### Creating subscription role bindings

Use the `subscriptionrolebinding` subcommand to assign a role binding to an account at subscription level, passing a valid role id.

##### Examples

* Create a role binding to an application `ocm support create applicationrolebinding [accountID] [roleID]`
* Create a role binding to an organization `ocm support create organizationrolebinding [accountID] [orgID] [roleID]`
* Create a role binding to a subscription `ocm support create subscriptionrolebinding [accountID] [subscriptionID] [roleID]`

### Delete

The `delete` command deletes the given resource.

#### Deleting labels

##### Deleting account labels

Use the `accountlabel` subcommand to delete a label from an account, passing the label key. 

##### Deleting organization labels

Use the `organizationlabel` subcommand to delete a label from an organization, passing the label key. 

##### Deleting subscription labels

Use the `subscriptionlabel` subcommand to delete a label from a subscription, passing the label key. 

##### Examples

* Delete a label from an account `ocm support delete accountlabel [accountID] [key]`
* Delete a label from an organization `ocm support delete organizationlabel [orgID] [key]`
* Delete a label from a subscription `ocm support delete subscriptionlabel [subscriptionID] [key]`

#### Deleting capabilities

##### Deleting account capabilities

Use the `accountcapability` subcommand to delete a capability from an account, passing the valid capability key. 

##### Deleting organization capabilities

Use the `organizationcapability` subcommand to delete a capability from an organization, passing the valid capability key. 

##### Deleting subscription capabilities

Use the `subscriptioncapability` subcommand to delete a capability from a subscription, passing a valid capability key. 

##### Examples

* Delete a capability from an account `ocm support delete accountcapability [accountID] [capabilityKey]`
* Delete a capability from an organization `ocm support delete organizationcapability [orgID] [capabilityKey]`
* Delete a capability from a subscription `ocm support delete subscriptioncapability [subscriptionID] [capabilityKey]`

##### Deleting type independent capabilities

Use the `capabilities` subcommand to provide filter value to search matching capabilities and delete them. By default the dry run flag will be enabled. Set `dry-run` flag to false to actually remove the resource.

The following flags are available for `delete capability`:

```
--dry-run                    If false, it will execute the delete command call in instead of a dry run.
--max-records                Maximum number of affected records. Defaults to 100. Only effective when dry-run is set to false.
-h, --help                   help for create
```

##### Examples

* Delete all capabilities where key is 'capability.account.create_moa_clusters' (with no dry run) `ocm support delete capabilities "key = 'capability.account.create_moa_clusters'" --dry-run=false --max-records=1000`

##### Deleting type independent capability

Use the `capability` subcommand to delete a capability by passing the ID. By default the dry run flag will be enabled. Set `dry-run` flag to false to actually remove the resource.

The following flags are available for `delete capability`:

```
--dry-run                    If false, it will execute the delete command call in instead of a dry run.
-h, --help                   help for create
```

##### Examples

* Delete a capability by its ID `ocm support delete capability [capabilityID]`

#### Deleting registry credentials

Use the `registrycredentials` subcommand to to delete a specific registry credential, or all registry credentials, for the passed accountID.

The following flags are available for `registrycredentials delete`:

```
--all                        If true, deletes all registry credentials for the given account ID.
-h, --help                   help for delete
```

#### Examples

* Delete a specific registry credential for a specific account `ocm support delete registrycredentials [accountID] [registryCredentialID]`
* Delete all registry credentials for a specific account `ocm support delete registrycredentials [accountID] --all`

#### Deleting role bindings

##### Deleting application role bindings

Use the `applicationrolebinding` subcommand to remove a role binding from an account at application level, passing a valid role id. 

##### Creating organization role bindings

Use the `organizationrolebinding` subcommand to remove a role binding from an account at organization level, passing a valid role id.

##### Creating subscription role bindings

Use the `subscriptionrolebinding` subcommand to remove a role binding from an account at subscription level, passing a valid role id.

##### Examples

* For the given account, delete a role binding created at application level using the roleID `ocm support delete applicationrolebinding [accountID] [roleID]`
* For the given account, delete a role binding created at organization level using the roleID `ocm support delete organizationrolebinding [accountID] [orgID] [roleID]`
* For the given account, delete a role binding created at subscription level using the roleID `ocm support delete subscriptionrolebinding [accountID] [subscriptionID] [roleID]`


### Patch

The `patch` command patches the given resource.

#### Patching accounts

Use the `accounts` subcommand and provide filter value to search matching accounts and patch them. Pass the JSON body for the patch request in terminal using `echo '{<PATCH_BODY>}' | ` before the actual command. By default the dry run flag will be enabled. Pass `dry-run=false` flag to actually patch the resource.

The following flags are available for `patch accounts`:

```
--dry-run                    If false, it will execute the patch command call in instead of a dry run.
--max-records                Maximum number of affected records. Defaults to 100. Only effective when dry-run is set to false.
-h, --help                   help for patch
```

##### Examples

* Patch accounts and change the last name to 'Doe' for accounts with username ending with 'doe' (no dry run) `echo '{ "last_name": "Doe" }' | ocm support patch accs "username like '%doe'" --dry-run=false`

#### Patching account

Use the `account` subcommand to patch an account by passing the ID. Pass the JSON body for the patch request in terminal using `echo '{<PATCH_BODY>}' | ` before the actual command. By default the dry run flag will be enabled. Pass `dry-run=false` flag to actually patch the resource.

The following flags are available for `patch account`:

```
--dry-run                    If false, it will execute the patch command call in instead of a dry run.
-h, --help                   help for patch
```

##### Examples

* Patch an account by its ID and change the first name to 'John' (dry run) `echo '{ "first_name": "John" }' | ocm support patch account [accID]`

#### Patching organizations

Use the `organizations` subcommand and provide filter value to search matching organizations and patch them. Pass the JSON body for the patch request in terminal using `echo '{<PATCH_BODY>}' | ` before the actual command. By default, the dry run flag will be enabled. Pass `dry-run=false` flag to actually patch the resource.

The following flags are available for `patch organizations`:

```
--dry-run                    If false, it will execute the patch command call in instead of a dry run.
--max-records                Maximum number of affected records. Defaults to 100. Only effective when dry-run is set to false.
-h, --help                   help for patch
```

##### Examples

* Patch all organizations with names starting with "Red Hat" and change the name to "Red Hat Inc." (no dry run and set maxRecords more than the actual number of affected records) `echo '{ "name": "Red Hat Inc." }' | ocm support patch orgs "name like 'Red Hat%' --dry-run=false --max-records=1000`


#### Patching organization

Use the `organization` subcommand to patch an organization by passing the ID. Pass the JSON body for the patch request in terminal using `echo '{<PATCH_BODY>}' | ` before the actual command. By default, the dry run flag will be enabled. Pass `dry-run=false` flag to actually patch the resource.

The following flags are available for `patch organizations`:

```
--dry-run                    If false, it will execute the patch command call in instead of a dry run.
-h, --help                   help for patch
```

##### Examples

* Patch an organization by its ID and change the externalID (dry run) `echo '{ "external_id": "12541229" }' | ocm support patch organization [orgID]`

#### Patching subscriptions

Use the `subscriptions` subcommand and provide filter value to search matching subscriptions and patch them. Pass the JSON body for the patch request in terminal using `echo '{<PATCH_BODY>}' | ` before the actual command. By default the dry run flag will be enabled. Pass `dry-run=false` flag to actually patch the resource.

The following flags are available for `patch subscriptions`:

```
--dry-run                    If false, it will execute the patch command call in instead of a dry run.
--max-records                Maximum number of affected records. Defaults to 100. Only effective when dry-run is set to false.
-h, --help                   help for patch
```

##### Examples

* Patch subscriptions and change the support level to Self-Support for subscriptions with 'Reserved' status (no dry run) `echo '{ "support_level": "Self-Support" }' | ocm support patch subs "status='Reserved'" --dry-run=false`
* Patch all subscriptions of an organization and change the status to Archived (no dry run and set maxRecords more than the actual number of affected records) `echo '{ "status": "Archived" }' | ocm support patch subs "organization_id='[orgID]' --dry-run=false --max-records=1000`

#### Patching subscription

Use the `subscription` subcommand to patch a subscription by passing the ID. Pass the JSON body for the patch request in terminal using `echo '{<PATCH_BODY>}' | ` before the actual command. By default the dry run flag will be enabled. Pass `dry-run=false` flag to actually patch the resource.

The following flags are available for `patch subscription`:

```
--dry-run                    If false, it will execute the patch command call in instead of a dry run.
-h, --help                   help for patch
```

##### Examples

* Patch a subscription by its ID and change the status to 'Reserved' (dry run) `echo '{ "status": "Reserved" }' | ocm support patch subscriptions [subID]`

#### Deleting an account

Use the `account` subcommand to delete a specific account for the passed accountID.

The following flags are available for `account delete`:

```
--dry-run                    If false, deletes the account for the given account ID, defaults to true.
-h, --help                   help for delete
```

#### Examples

* Delete a specific account `ocm support delete account [accountID]`
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

### Create

The `create` command creates the given resource.

#### Creating account labels

Use the `accountLabel` subcommand to assign a label to an account, passing a key and a value. 

The following flags are available for `create accountLabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

#### Creating organization labels

Use the `organizationLabel` subcommand to assign a label to an organization, passing a key and a value. 

The following flags are available for `create organizationLabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

#### Creating subscription labels

Use the `subscriptionLabel` subcommand to assign a label to a subscription, passing a key and a value. 

The following flags are available for `create subscriptionLabel`:

```
--external                   If true, sets the internal flag for labels as false.
-h, --help                   help for create
```

#### Creating account capabilities

Use the `accountCapability` subcommand to assign a capability to an account, passing a valid capability key. 

#### Creating organization capabilities

Use the `organizationCapability` subcommand to assign a capability to an organization, passing a valid capability key. 

#### Creating subscription capabilities

Use the `subscriptionCapability` subcommand to assign a capability to a subscription, passing a valid capability key. 

#### Creating registry credentials

Use the `registryCredentials` subcommand to create registry credentials for current account. 

#### Examples

* Create a label to an account `ocm support create accountLabel [accountID] [key] [value]`
* Create a label to an organization `ocm support create organizationLabel [orgID] [key] [value]`
* Create a label to a subscription (with internal flag as false) `ocm support create subscriptionLabel [subscriptionID] [key] [value] --external`
* Create a capability to an account `ocm support create accountCapability [accountID] [capabilityKey]`
* Create a capability to an organization `ocm support create organizationCapability [orgID] [capabilityKey]`
* Create a capability to a subscription `ocm support create subscriptionCapability [subscriptionID] [capabilityKey]`
* Create a capability to a subscription `ocm support create subscriptionCapability [subscriptionID] [capabilityKey]`
* Create registryCredentials `ocm support create registryCredentials`

### Delete

The `delete` command deletes the given resource.

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

### Find

The `find` command finds the given resource.

#### Finding an account

Use the `accounts` subcommand to find one or more accounts, passing as argument one of the following:

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

#### Finding an organization

Use the `organizations` subcommand to find one or more organizations, passing as argument one of the following:

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

#### Finding registry credentials

Use the `registryCredentials` subcommand to to find registry credentials, passing accountID.

#### Examples

* Find an account by email `ocm support find accounts user@redhat.com`
* Find an account and include its roles in the results `ocm support find accounts [accountID] --fetchRoles`
* Find an account and include its labels in the results `ocm support find accounts [accountID] --fetchLabels`
* Find an account and include its capabilities in the results `ocm support find accounts [accountID] --fetchCapabilities`
* Find all accounts from an organization `ocm support find accounts [organizationID] --all`
* Find an organization by its externalID: `ocm support find organizations [organizationExternalID]`
* Find an organization and include its subscriptions: `ocm support find organizations [organizationID] --fetchSubscriptions`
* Find an organization and include its labels `ocm support find organizations [organizationID] --fetchLabels --fetchCapabilities`
* Find an organization and include its capabilities `ocm support find organizations [organizationID] --fetchCapabilities`
* Find an organization and include its quota: `ocm support find organizations [organizationID] --fetchQuota`
* Show registry credentials for a specific account `ocm support find registryCredentials [accountID]`
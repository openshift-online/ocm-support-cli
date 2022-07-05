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

To see all available commands, run `ocm support version -h`.

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
-h, --help                   help for find
```

#### Examples

* Find an account by email `ocm support accounts find user@redhat.com`
* Find an account and include its roles in the results `ocm support accounts find [accountID] --fetchRoles`
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
-h, --help             help for find
```

#### Examples

* Find an organization by its externalID: `ocm support organizations find [organizationExternalID]`
* Find an organization and include its subscriptions: `ocm support organizations find [organizationID] --fetchSubscriptions`
* Find an organization and include its quota: `ocm support organizations find [organizationID] --fetchQuota`

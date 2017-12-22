// Code generated by stringdata. DO NOT EDIT.

package schema

// SiteSchemaJSON is the content of the file "site.schema.json".
const SiteSchemaJSON = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "https://sourcegraph.com/v1/site.schema.json#",
  "title": "Site configuration",
  "description": "Configuration for a Sourcegraph Server site.",
  "type": "object",
  "additionalProperties": false,
  "properties": {
    "auth.userOrgMap": {
      "description":
        "Ensure that matching users are members of the specified orgs (auto-joining users to the orgs if they are not already a member). Provide a JSON object of the form <tt>{\"*\": [\"org1\", \"org2\"]}</tt>, where org1 and org2 are orgs that all users are automatically joined to. Currently the only supported key is <tt>\"*\"</tt>.",
      "type": "object",
      "additionalProperties": {
        "type": "array",
        "items": {
          "type": "string"
        }
      }
    },
    "appID": {
      "description":
        "Application ID to attribute front end user logs to. Providing this value will send usage data back to Sourcegraph (no private code is sent and URLs are sanitized to prevent leakage of private data).",
      "type": "string"
    },
    "appURL": {
      "description": "Publicly accessible URL to web app (e.g., what you type into your browser).",
      "type": "string"
    },
    "disableTelemetry": {
      "description":
        "Prevent usage data from being sent back to Sourcegraph (no private code is sent and URLs are sanitized to prevent leakage of private data).",
      "type": "boolean"
    },
    "tlsCert": {
      "description": "TLS certificate for the web app.",
      "type": "string"
    },
    "tlsKey": {
      "description": "TLS key for the web app.",
      "type": "string"
    },
    "corsOrigin": {
      "description": "Value for the Access-Control-Allow-Origin header returned with all requests.",
      "type": "string"
    },
    "autoRepoAdd": {
      "description": "Automatically add external public repositories on demand when visited.",
      "type": "boolean"
    },
    "disablePublicRepoRedirects": {
      "description":
        "Disable redirects to sourcegraph.com when visiting public repositories that can't exist on this server.",
      "type": "boolean"
    },
    "phabricator": {
      "description":
        "JSON array of configuration for Phabricator hosts. See Phabricator Configuration section for more information.",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "url": {
            "description": "URL of a Phabricator instance, such as https://phabricator.example.com",
            "type": "string"
          },
          "token": {
            "description": "API token for the Phabricator instance.",
            "type": "string"
          },
          "repos": {
            "description": "The list of repos available on Phabricator.",
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "path": {
                  "description": "Display path for the url e.g. gitolite/my/repo",
                  "type": "string"
                },
                "callsign": {
                  "description": "The unique Phabricator identifier for the repo, like 'MUX'.",
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "phabricatorURL": {
      "description": "(Deprecated: Use Phabricator) URL of Phabricator instance.",
      "type": "string"
    },
    "github": {
      "description":
        "JSON array of configuration for GitHub hosts. See GitHub Configuration section for more information.",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "url": {
            "description":
              "URL of a GitHub instance, such as https://github.com or https://github-enterprise.example.com",
            "type": "string"
          },
          "token": {
            "description": "A GitHub personal access token with repo and org scope.",
            "type": "string"
          },
          "certificate": {
            "description": "TLS certificate of a GitHub Enterprise instance.",
            "type": "string"
          },
          "repos": {
            "description": "Optional list of additional repositories to clone (in \"owner/repo\" format).",
            "type": "array",
            "items": { "type": "string" }
          }
        }
      }
    },
    "githubClientID": {
      "description": "Client ID for GitHub.",
      "type": "string"
    },
    "githubClientSecret": {
      "description": "Client secret for GitHub.",
      "type": "string"
    },
    "githubPersonalAccessToken": {
      "description": "(Deprecated: Use GitHub) Personal access token for GitHub. ",
      "type": "string"
    },
    "githubEnterpriseURL": {
      "description": "(Deprecated: Use GitHub) URL of GitHub Enterprise instance from which to sync repositories.",
      "type": "string"
    },
    "githubEnterpriseCert": {
      "description":
        "(Deprecated: Use GitHub) TLS certificate of GitHub Enterprise instance, if from a CA that's not part of the standard certificate chain.",
      "type": "string"
    },
    "githubEnterpriseAccessToken": {
      "description": "(Deprecated: Use GitHub) Access token to authenticate to GitHub Enterprise API.",
      "type": "string"
    },
    "gitoliteHosts": {
      "description": "Space separated list of mappings from repo name prefix to gitolite hosts.",
      "type": "string"
    },
    "gitOriginMap": {
      "description":
        "Space separated list of mappings from repo name prefix to origin url, for example \"github.com/!https://github.com/%.git\".",
      "type": "string"
    },
    "repos.list": {
      "description": "JSON array of configuration for external repositories.",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "type": {
            "description": "Type of the version control system for this repository, such as \"git\"",
            "type": "string",
            "enum": ["git"],
            "default": "git"
          },
          "url": {
            "description": "Clone URL for the repository, such as git@example.com:my/repo.git",
            "type": "string"
          },
          "path": {
            "description": "Display path on Sourcegraph for the repository, such as my/repo",
            "type": "string"
          }
        }
      }
    },
    "inactiveRepos": {
      "description": "Comma-separated list of repos to consider 'inactive' (e.g. while searching).",
      "type": "string"
    },
    "lightstepAccessToken": {
      "description": "Access token for sending traces to LightStep.",
      "type": "string"
    },
    "lightstepProject": {
      "description": "The project id on LightStep, only used for creating links to traces.",
      "type": "string"
    },
    "noGoGetDomains": {
      "description": "List of domains to NOT perform go get on. Separated by ','.",
      "type": "string"
    },
    "repoListUpdateInterval": {
      "description": "Interval (in minutes) for checking code hosts (e.g. gitolite) for new repositories.",
      "type": "integer"
    },
    "ssoUserHeader": {
      "description": "Header injected by an SSO proxy to indicate the logged in user.",
      "type": "string"
    },
    "oidcProvider": {
      "description": "The URL of the OpenID Connect Provider",
      "type": "string"
    },
    "oidcClientID": { "description": "OIDC Client ID", "type": "string" },
    "oidcClientSecret": {
      "description": "OIDC Client Secret",
      "type": "string"
    },
    "oidcEmailDomain": {
      "description": "Whitelisted email domain for logins, e.g. 'mycompany.com'",
      "type": "string"
    },
    "oidcOverrideToken": {
      "description": "Token to circumvent OIDC layer (testing only)",
      "type": "string"
    },
    "samlIDProviderMetadataURL": {
      "description": "SAML Identity Provider metadata URL (for dyanmic configuration of SAML Service Provider)",
      "type": "string"
    },
    "samlSPCert": {
      "description": "SAML Service Provider certificate",
      "type": "string"
    },
    "samlSPKey": {
      "description": "SAML Service Provider private key",
      "type": "string"
    },
    "searchScopes": {
      "description":
        "JSON array of custom search scopes (e.g., [{\"name\":\"Text Files\",\"value\":\"file:\\.txt$\"}])",
      "type": "array",
      "items": {
        "type": "object",
        "$ref": "settings.schema.json#/definitions/SearchScope"
      }
    },
    "htmlHeadTop": {
      "description": "HTML to inject at the top of the <head> element on each page, for analytics scripts",
      "type": "string"
    },
    "htmlHeadBottom": {
      "description": "HTML to inject at the bottom of the <head> element on each page, for analytics scripts",
      "type": "string"
    },
    "htmlBodyTop": {
      "description": "HTML to inject at the top of the <body> element on each page, for analytics scripts",
      "type": "string"
    },
    "htmlBodyBottom": {
      "description": "HTML to inject at the bottom of the <body> element on each page, for analytics scripts",
      "type": "string"
    },
    "licenseKey": {
      "description": "License key. You must purchase a license to obtain this.",
      "type": "string"
    },
    "mandrillKey": {
      "description": "Key for sending mails via Mandrill.",
      "type": "string"
    },
    "maxReposToSearch": {
      "description":
        "The maximum number of repos to search across. The user is prompted to narrow their query if exceeded. The value 0 means unlimited.",
      "type": "integer"
    },
    "adminUsernames": {
      "description": "Space-separated list of usernames that indicates which users will be treated as instance admins",
      "type": "string"
    },
    "executeGradleOriginalRootPaths": {
      "description":
        "Java: A comma-delimited list of patterns that selects repository revisions for which to execute Gradle scripts, rather than extracting Gradle metadata statically. **Security note:** these should be restricted to repositories within your own organization. A percent sign ('%') can be used to prefix-match. For example, <tt>git://my.internal.host/org1/%,git://my.internal.host/org2/repoA?%</tt> would select all revisions of all repositories in org1 and all revisions of repoA in org2.",
      "type": "string"
    },
    "privateArtifactRepoID": {
      "description":
        "Java: Private artifact repository ID in your build files. If you do not explicitly include the private artifact repository, then set this to some unique string (e.g,. \"my-repository\").",
      "type": "string"
    },
    "privateArtifactRepoURL": {
      "description":
        "Java: The URL that corresponds to privateArtifactRepoID (e.g., http://my.artifactory.local/artifactory/root).",
      "type": "string"
    },
    "privateArtifactRepoUsername": {
      "description": "Java: The username to authenticate to the private Artifactory.",
      "type": "string"
    },
    "privateArtifactRepoPassword": {
      "description": "Java: The password to authenticate to the private Artifactory.",
      "type": "string"
    },
    "settings": {
      "description": "Site settings. Organization and user settings override site settings.",
      "type": "object",
      "$ref": "https://sourcegraph.com/v1/settings.schema.json#"
    }
  }
}
`

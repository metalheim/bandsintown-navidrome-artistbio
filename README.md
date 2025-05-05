# bandsintown-navidrome-artistbio
Navidrome plugin that fetches band tour dates and displays them as artist bio.

This is my first project in golang and 

## Usage
Add the plugin files to your plugin folder.
The plugin folder structure should look like this:
```
plugins/
├── bandsintown-navidrome-artistbio/
│   ├── plugin.wasm         
│   └── manifest.json     
├── another-plugin/
│   ├── plugin.wasm
│   └── manifest.json
```
If this is your first navidrome plugin, enable plugins in general via your navidrome configuration file 
``` toml
[Plugins]
Enabled = true
Folder = "/path/to/plugins"
```
Also add the plugin-specific configuration (api_key and app_id)
``` toml
[PluginConfig.bandsintown-navidrome-artistbio]
api_key = "4bfaeac7-557d-4253-987f-f35160aafcf9"
app_id = "12345567"
```
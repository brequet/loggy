# Todo List
================

- [ ] Parser
    - [ ] When parsing a file, determine log format one time instead of each time a log is parsed
    - [ ] Activemq logs does not have a timezone -> read local timezone and adjust logs accordingly => reflect it in logFormat conf ?
    - [x] If conf file given contains paths to log folders, use them (+ use default conf file for my project)
    - [x] For ingest, add a flag to ignore conf log dir path, instead parse directory given as argument (log_dir/app_name/*.log)
- [x] Embed frontend into Go app
- [ ] Add logger to SQLite module
- [x] Move server instantiation into calling CLI command
- [ ] Handle parsing using configuration file
- [x] Check if DB exists at serve command
- [x] Fix bug where if log scrolled and then apply filter, several page are fetched at the same time (instead of one), and the page are coming in wrong order (so the logs are not in order)
- [ ] Add a left panel to show log count by app and by level
- [x] Make configuration file importable (override default one)
- [ ] For each log entry (table row), make it globally clickable -> open menu
    - Set filters
        - App
            - Include/exclude
        - Level
            - Include/exclude
        - Dates (start/end)
- [ ] At end of ingest, log some stats
# Todo List
================

- [x] Embed frontend into Go app
- [ ] Add logger to SQLite module
- [x] Move server instantiation into calling CLI command
- [ ] Handle parsing using configuration file
- [x] Check if DB exists at serve command
- [x] Fix bug where if log scrolled and then apply filter, several page are fetched at the same time (instead of one), and the page are coming in wrong order (so the logs are not in order)
- [ ] Add a left panel to show log count by app and by level
- [ ] Make configuration file importable (override default one)

# Netrunner Card Issue Importer

A command line tool to import a CSV file of new Netrunner cards as Github Issues.
Each entry in the CSV file will generate a new Github Issue in the specified repo with the card name as the title and the card description as the body. A label of `New Card` will be set on the Issue.

# Usage
```
% go build cmd
% ./import -token "github ouath token" -repo "user/repo" new-cards.csv
```

# CSV Format
The CSV file is expected to have a header as the first line and will be skipped.
The CSV file expects to have the following fields:
```
advancement-requirement;agenda-points;base-link;cost;deck-limit;faction;id;influence-limit;influence-value;memory-cost;minimum-deck-size;position;quantity;side;strength;subtype;text;title;trash-cost;type;uniqueness
```


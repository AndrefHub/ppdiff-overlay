## Installation tutorial

### 1. Get Client ID and Client Secret from osu!
1.1. Go to [this page](https://osu.ppy.sh/home/account/edit)
1.2. Scroll down to __OAuth__ and click _New OAuth Application_
1.3. Enter whatever name you want and click __Register Application__
1.4. Copy your new application's ID and Secret and paste them in config.ini

### 2. Add new widgets
2.1. All widgets should be located in /static folder so this application can access them
2.2. Widget should be located in it's own folder and have index.html as it's main page
Project's tree example: 
.
├── ppdiff-overlay.exe
├── config.ini
└── static
    └── dummy-widget
        ├── index.js
        ├── index.css
        └── index.html

### 3. Run executable and enjoy
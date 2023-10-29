const config = {
  plugin: ["prettier-plugin-go-template"],
  overrides: [
    {
      files: ["*.html"],
      options: {
        parser: "go-template",
      },
    },
  ],
};

module.exports = config;

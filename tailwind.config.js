module.exports = {
  purge: ["./src/**/*.html", "./src/**/*.vue"],
  theme: {
    inset: {
      "0": 0,
      "1": "0.25rem",
      "2": "0.5rem",
      "3": "0.75rem",
      "4": "1rem",
      "-2": "-0.5rem",
      "-07": "-0.7rem",
      "-3": "-0.75rem",
      "-4": "-1rem",
      "-8": "-2rem",
      "-16": "-4rem"
    }
  },
  variants: {
    borderColor: ["responsive", "hover", "focus", "focus-within"]
  },
  plugins: []
};

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./App.{js,jsx,ts,tsx}", "./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: "#4F46E5",
        secondary: "#10B981",
        highlight: "#F59E0B",
        background: "#F3F4F6",
        surface: "#FFFFFF",
        textPrimary: "#111827",
        textSecondary: "#6B7280",
      }
    },
  },
  plugins: [],
};

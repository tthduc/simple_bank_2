/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

// eslint that will help you to analyze your code and find errors
module.exports = {
  root: true,
  'extends': [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier/skip-formatting'
  ],
  parserOptions: {
    ecmaVersion: 'latest'
  }
}

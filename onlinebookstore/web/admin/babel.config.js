module.exports = {
  
  presets: [
    '@vue/cli-plugin-babel/preset',
    
  ],
  plugins: [
      [
       "import",
       { libraryName: "ant-design-vue", libraryDirectory: "es", style: "css" }
       ],
    ],
  presets: [ [ "@vue/app", { useBuiltIns: "entry" } ] ],
  
}

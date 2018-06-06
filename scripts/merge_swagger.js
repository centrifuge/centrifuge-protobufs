const swaggermerge = require('swagger-merge')
const fs = require('fs')
const path = require('path')
const util = require('util');
const exec = util.promisify(require('child_process').exec);

const swaggerFilesPath = path.resolve(__dirname, '../gen/swagger')
const swaggerJsonPath = path.resolve(__dirname, '../gen/swagger.json')
const swaggerHtmlPath = path.resolve(__dirname, '../gen/swagger/html')
const info = {
    version: "0.0.1",
    title: "Centrifuge OS Node API",
    description: "\n"
}
const schemes = ['https']
const host = "localhost"
const pathPrefix = ""

// From: https://gist.github.com/kethinov/6658166
var getSwaggerFiles = function(dir, filelist) {
    var files = fs.readdirSync(dir);
    filelist = filelist || [];
    files.forEach(function(file) {
        if (fs.statSync(dir + '/' + file).isDirectory()) {
            getSwaggerFiles(dir + '/' + file, filelist);
        }
        else {
            if (file.indexOf(".swagger.json") > 0) {
                filelist.push(path.join(dir, file));
            }
        }
    });
    return filelist;
};

var files = getSwaggerFiles(swaggerFilesPath);
swaggerModules = []
files.forEach(function (f) {
    swaggerModules.push(require(f))
})

swaggermerge.on('warn', function (msg) {
    console.log(msg)
})
swaggermerge.on('err', function (msg) {
    console.error(msg)
    process.exit(1)
})

var merged = swaggermerge.merge(swaggerModules, info, pathPrefix, host, schemes)
var json = JSON.stringify(merged);
console.log("Merged swagger.json, writing to:", swaggerJsonPath);
fs.writeFileSync(swaggerJsonPath, json);

var spectacles = exec('spectacle gen/swagger.json -t '+swaggerHtmlPath).then(function (msg) {
    console.log(msg['stdout'])
    console.log("Wrote html files to: ", swaggerHtmlPath)
    process.exit(0)
}).catch(function (err) {
    console.error(err)
    process.exit(1)
});


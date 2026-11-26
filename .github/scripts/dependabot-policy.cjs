// Fail closed: only Go patch updates with unchanged toolchain requirements qualify.
function directives(source) {
  return source.split('\n').map(line => line.replace(/\/\/.*$/, '').trim())
    .filter(line => /^(go|toolchain)\s/.test(line)).join('\n');
}
function eligible({ecosystem, updateType, files, before, after}) {
  return ecosystem === 'gomod' && updateType === 'version-update:semver-patch' &&
    files.length > 0 && files.every(file => ['go.mod', 'go.sum'].includes(file)) &&
    /^go\s/m.test(before) && /^go\s/m.test(after) && directives(before) === directives(after);
}
module.exports = {eligible};

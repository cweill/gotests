const {test} = require('node:test');
const assert = require('node:assert/strict');
const {eligible} = require('./dependabot-policy.cjs');
const base = {ecosystem: 'gomod', updateType: 'version-update:semver-patch', files: ['go.mod', 'go.sum'],
  before: 'module example\n\ngo 1.24.0\ntoolchain go1.24.5\nrequire example.com/a v1.0.0\n',
  after: 'module example\n\ngo 1.24.0\ntoolchain go1.24.5\nrequire example.com/a v1.0.1\n'};
test('allows Go patch updates', () => assert.equal(eligible(base), true));
for (const [name, change] of Object.entries({
  minor: {updateType: 'version-update:semver-minor'},
  actions: {ecosystem: 'github-actions'},
  source: {files: ['go.mod', 'gotests.go']},
  empty: {files: []},
  goVersion: {after: base.after.replace('go 1.24.0', 'go 1.25.0')},
  toolchain: {after: base.after.replace('go1.24.5', 'go1.25.0')},
  removedToolchain: {after: base.after.replace('toolchain go1.24.5\n', '')},
  missingModule: {after: ''},
})) test(`rejects ${name}`, () => assert.equal(eligible({...base, ...change}), false));

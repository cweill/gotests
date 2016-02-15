import sublime, sublime_plugin
import subprocess

class gotestsCommand(sublime_plugin.TextCommand):
	def run(self, edit):
		fn = self.view.file_name()
		if fn and fn.endswith('.go') and not fn.endswith('_test.go'):
			settings = sublime.load_settings("GoTests.sublime-settings")
			fs = []
			for s in self.view.sel():
				line = self.function_line(s.begin())
				i = line.begin()
				while i <= s.end():
					f = self.function_name(line)
					i = line.end() + 1
					line = self.view.line(i)
					if not f:
						continue
					fs.append(f)
			try:
				gotests = settings.get("gotests_cmd", "gotests")
				cmd = [gotests, '-w', '-only=^(' + "|".join(fs) + ')$', fn]
				proc = subprocess.Popen(cmd, stdout=subprocess.PIPE)
				print(proc.stdout.read().decode("utf-8").replace('\r\n', '\n'))
			except OSError as e:
				sublime.message_dialog("GoTests error: " + str(e))
				return False
			return True
		return False

	# Returns a function signature's line from a point in its body.
	def function_line(self, point):
		line = self.view.line(point)
		if self.is_blank(line):
			return line
		above = line
		while not self.is_blank(above) and not self.function_name(above) and above.begin() > 0:
			above = self.view.line(above.begin() - 1)
		return above

	# Return whether the line is blank.
	def is_blank(self, line):
		return self.view.substr(line.begin()) == '\n'

	# Returns the name of the function if the given line is a method signature.
	def function_name(self, line):
		if self.view.substr(self.view.word(line.begin())) != "func":
			return None
		i = line.begin()
		while i <= line.end():
			word = self.view.word(i)
			i = word.end() + 1
			c = self.view.substr(word.end())
			if c == "(":
				return self.view.substr(word)
		return None

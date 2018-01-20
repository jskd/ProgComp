package prog.comp2018.scableur.utils

import util.matching.Regex

object RegexUtils {
  class RichRegex(self: Regex) {
    def =~(s: String) = self.pattern.matcher(s).matches
  }
  implicit def regexToRichRegex(r: Regex) = new RichRegex(r)
}



package presume

import scalatags.Text.all.*

opaque type Classes = List[String]

object Classes:
  def apply(from: String*): Classes = from.toList

  extension (c: Classes)
    def toList: List[String] = c
    def toAttrPair: AttrPair = `class` := c.mkString(" ")
    def ++(other: Classes): Classes =
      c ++ other
    def ::(other: String): Classes =
      other :: c
    def :+(other: String): Classes =
      c :+ other
    def +:(other: String): Classes =
      other +: c

  given Conversion[Classes, AttrPair] = c => c.toAttrPair

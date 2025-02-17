package presume

import cats.data.*
import cats.syntax.all.*
import com.monovore.decline.*
import com.monovore.decline.effect.*
import cats.effect.*
import scala.util.Try
import presume.models.*
import presume.models.syntax.given
import smithy4s.Blob
import smithy4s.xml.Xml

object Presume extends CommandIOApp(
  "presume",
  "Generate my resume!",
  true,
  "0.0.1"
):
  val resumeContentDecoder = Xml.decoders.fromSchema(ResumeContent.schema)

  override def main: Opts[IO[ExitCode]] =
    (
      Opts.argument[os.Path]("input").mapValidated(inputPath =>
        Either.catchNonFatal(os.read(inputPath))
          .leftMap(_.getMessage)
          .map(Blob.apply)
          .flatMap(resumeContentDecoder.decode(_).leftMap(_.getMessage))
          .toValidatedNel
          .map(inputPath -> _)
      ),
      Opts.option[os.Path]("output", "The output file (defaults to resume name with .html", "o").orNone
    ).mapN: (input, output) =>
      IO.println(show"Input is ${input._2}")
        .flatMap(_ => IO.println(s"output is $output"))
        .flatMap(_ => IO.blocking(
          os.write.over(
            output.getOrElse(os.pwd / (input._1.baseName + ".html")),
            ResumeTemplate(input._2).toString)
          )
        )
        .as(ExitCode.Success)

  given Argument[os.Path] = Argument.from("path"): value =>
    Either.catchNonFatal(os.Path(value))
      .leftMap(_.getMessage)
      .recoverWith:
        case pathError =>
          Either.catchNonFatal(os.RelPath(value))
            .map(os.pwd / _)
            .leftMap(_.getMessage)
            .adaptError:
              case relPathError =>
                s"Given path is neither relative nor absolute. Errors are ['$pathError', '$relPathError']"
      .toValidatedNel


package presume

import cats.*
import cats.syntax.all.*
import presume.models.*
import presume.models.syntax.{*, given}
import scalatags.Text.all.*

object ResumeTemplate:

  val border = Classes("border-1", "rounded-sm", "border-indigo-300", "border-solid")
  val sectionShadow = Classes("shadow-lg", "shadow-indigo-500/20")
  val sectionClasses =
    Classes("mt-4", "ml-4", "mr-4", "p-2", "shadow-lg") ++
      sectionShadow ++
      border

  def apply(content: ResumeContent) = html(
    head(
      meta(charset := "UTF-8"),
      script(src := "https://unpkg.com/@tailwindcss/browser@4")
    ),
    body(
      div(
        Classes("w-1/3", "mt-4", "ml-16", "p-2", "shadow-lg") ++
          sectionShadow ++
          border,
        p(
          Classes("text-2xl"),
          content.header.name,
        ) ::
        p(
          Classes("text-sm"),
          content.header.email
        ) ::
        content.header.gitHubName
          .map: name =>
            p(
              `class` := "text-sm",
              name
            )
          .toList
      ),
      div(
        sectionClasses,
        p(
          Classes("text-2xl"),
          "Experience"
        ),
        content.experience.map: experience =>
          div(
            Classes("mt-2"),
            p(
              Classes("text-xl"),
              experience.title
            ),
            p(
              Classes("text-md"),
              experience.business
            ),
            p(
              Classes("text-sm"),
              show"${experience.startDate} - ${experience.endDate}"
            ),
            ul(
              Classes("list-disc", "list-inside"),
              experience.experience.map: item =>
                li(
                  Classes("text-md"),
                  item.value
                )
            )
          )
      ),
      div(
        sectionClasses,
        p(
          Classes("text-2xl"),
          "Education"
        ),
        content.education.map: education =>
          div(
            p(
              Classes("text-lg"),
              education.schoolName
            ),
            p(
              Classes("text-sm"),
              show"${education.startDate} - ${education.endDate}"
            ),
            p(
              Classes("text-md"),
              education.certificateType
            ),
            ul(
              Classes("list-disc", "list-inside"),
              education.additional.map: add =>
                li(
                  Classes("text-md"),
                  add
                )
            )
          )
      ),
      div(
        sectionClasses,
        p(
          Classes("text-2xl"),
          "Skills"
        )
      )
    )
  )

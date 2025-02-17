package presume.models.syntax

import cats.*
import cats.syntax.all.*
import presume.models.*

given Show[Month] = Show.show: month =>
  val str = month.toString()
  str.head + str.tail.toLowerCase

given Show[Date] = Show.show: date =>
  show"${date.month} ${date.year}"

given Show[DateOrPresent] = Show.show:
  case DateOrPresent.PresentCase => "Present"
  case DateOrPresent.DateCase(date) => date.show

given Show[Header] = Show.show:
  case Header(name, email, gitHubName) =>
    show"Header(name: $name, email: $email, gitHubName: $gitHubName)"

given Show[ExperienceBullet] = Show.show:
  case ExperienceBullet(value) => value.trim()

given Show[ExperienceRange] = Show.show:
  case ExperienceRange(business, title, startDate, endDate, experience) =>
    show"ExperienceRange(business: $business, title: $title, startDate: $startDate, endDate: $endDate, experience: $experience)"

given Show[EducationItem] = Show.show:
  case EducationItem(schoolName, certificateType, additional, startDate, endDate) =>
    show"EducationItem(schoolName: $schoolName, certificateType: $certificateType, additional: $additional, startDate: $startDate, endDate: $endDate)"

given Show[ResumeContent] = Show.show:
  case ResumeContent(header, experience, education) =>
    show"ResumeContent(header: $header, experience: $experience, education: $education)"

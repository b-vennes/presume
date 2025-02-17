$version: "2"

namespace presume.models

enum Month {
  JANUARY
  FEBRUARY
  MARCH
  APRIL
  MAY
  JUNE
  JULY
  AUGUST
  SEPTEMBER
  OCTOBER
  NOVEMBER
  DECEMBER
}

structure Date {
  @required
  year: Integer

  @required
  month: Month
}

union DateOrPresent {
  date: Date
  present: Unit
}

string ExperienceBullet

list ExperienceList {
  member: ExperienceBullet
}

structure ExperienceRange {
  @required
  business: String

  @required
  title: String

  @required
  startDate: Date

  @required
  endDate: DateOrPresent

  @required
  experience: ExperienceList
}

list ExperienceSection {
  member: ExperienceRange
}

list AdditionalCertificates {
  member: String
}

structure EducationItem {
  @required
  schoolName: String

  @required
  certificateType: String

  @required
  additional: AdditionalCertificates

  @required
  startDate: Date

  @required
  endDate: DateOrPresent
}

list EducationSection {
  member: EducationItem
}

structure Header {
  @required
  name: String

  @required
  email: String

  gitHubName: String
}

structure ResumeContent {
  @required
  header: Header

  @required
  experience: ExperienceSection

  @required
  education: EducationSection
}

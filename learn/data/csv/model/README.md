# [CSV Data Forms](https://canvas.instructure.com/doc/api/file.sis_csv.html)

## users.csv

| Field Name                 | Data Type       | Required | Sticky | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| -------------------------- | --------------- | -------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| user_id                    | text            | ✓        |        | A unique identifier used to reference users in the enrollments table. This identifier must not change for the user, and must be globally unique. In the user interface, this is called the _SIS ID_.                                                                                                                                                                                                                                                                 |
| integration_id             | text            |          |        | A secondary unique identifier useful for more complex SIS integrations. This identifier must not change for the user, and must be globally unique.                                                                                                                                                                                                                                                                                                                   |
| login_id                   | text            | ✓        | ✓      | The name that a user will use to login to Instructure. If you have an authentication service configured _(like LDAP)_, this will be their username from the remote system.                                                                                                                                                                                                                                                                                           |
| password                   | text            |          |        | If the account is configured to use _LDAP_ or an _SSO protocol_ then this should not be set. Otherwise this is the password that will be used to login to Canvas along with the '**login_id**' above. Setting the password will in most cases log the user out of Canvas. The password can only be set one time. If the password has been set by the user or a previous sis import, it will not be changed.                                                          |
| ssha_password              | text            |          |        | Instead of a plain-text password, you can pass a pre-hashed password using the _SSHA password_ generation scheme in this field. While better than passing a plain text password, you should still encourage users to change their password after logging in for the first time.                                                                                                                                                                                      |
| authentication_provider_id | text or integer |          |        | The authentication provider this login is associated with. Logins associated with a specific provider can only be used with that provider. Legacy providers _(LDAP, CAS, SAML)_ will search for logins associated with them, or unassociated logins. New providers will only search for logins explicitly associated with them. This can be the _integer ID_ of the provider, or the type of the provider (in which case, it will find the first matching provider). |
| first_name                 | text            |          | ✓      | Given name of the user. If present, used to construct **full_name** and/or **sortable_name**.                                                                                                                                                                                                                                                                                                                                                                        |
| last_name                  | text            |          | ✓      | Last name of the user. If present, used to construct **full_name** and/or **sortable_name**.                                                                                                                                                                                                                                                                                                                                                                         |
| full_name                  | text            |          | ✓      | Full name of the user. Omit **first_name** and **last_name** if this is provided.                                                                                                                                                                                                                                                                                                                                                                                    |
| sortable_name              | text            |          | ✓      | Sortable name of the user. This is normally inferred from the user's name, but you can customize it here.                                                                                                                                                                                                                                                                                                                                                            |
| short_name                 | text            |          | ✓      | Display name of the user. This is normally inferred from the user's name, but you can customize it here.                                                                                                                                                                                                                                                                                                                                                             |
| email                      | text            |          |        | The email address of the user. This might be the same as **login_id** but would be used to set email for user and will tie the email to the login. It is recommended to omit this field over using fake email addresses for testing.                                                                                                                                                                                                                                 |
| pronouns                   | text            |          | ✓      | User's preferred pronouns. Can pass `"<delete>"` to remove the pronoun from the user.                                                                                                                                                                                                                                                                                                                                                                                |
| declared_user_type         | enum            |          |        | User's declared user type. Can be either `administrative`, `observer`, `staff`, `student`, `student_other`, or `teacher`. Can pass "<delete>" to remove the declared user type from the user.                                                                                                                                                                                                                                                                        |
| status                     | enum            | ✓        | ✓      | `active`, `suspended`, `deleted`                                                                                                                                                                                                                                                                                                                                                                                                                                     |

At least one form of name should be supplied. If a user is being created and no
name is given, the **login_id** will be used as the name.

When a user is `'deleted'` it will delete the login tied to the sis_id. If the
login is the last one, all of the users enrollments will also be deleted and
they won't be able to log in to the school's account. If you still want the
student to be able to log in but just not participate, leave the student
`'active'` but set the enrollments to 'completed'. If you want to leave a
student's enrollments intact, but not allow them to login, use the `'suspended'`
status.

## accounts.csv

| Field Name        | Data Type | Required | Sticky | Description                                                                                                                                                                                                                    |
| ----------------- | --------- | -------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| account_id        | text      | ✓        |        | A unique identifier used to reference accounts in the enrollments data. This identifier must not change for the account, and must be globally unique. In the user interface, this is called the _SIS ID_.                      |
| parent_account_id | text      | ✓        | ✓      | The account identifier of the parent account. If this is blank the parent account will be the root account. Note that even if all values are blank, the column must be included to differentiate the file from a group import. |
| name              | text      | ✓        | ✓      | The name of the account                                                                                                                                                                                                        |
| status            | enum      | ✓        |        | `active`, `deleted`                                                                                                                                                                                                            |
| integration_id    | text      |          |        | Sets the **integration_id** of the account                                                                                                                                                                                     |

Any account that will have child accounts must be listed in the csv before any
child account references it.

## terms.csv

| Field Name                    | Data Type | Required | Sticky | Description                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| ----------------------------- | --------- | -------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| term_id                       | text      | ✓        |        | A unique identifier used to reference terms in the enrollments data. This identifier must not change for the account, and must be globally unique. In the user interface, this is called the _SIS ID_.                                                                                                                                                                                                                                             |
| name                          | text      | ✓        | ✓      | The name of the term                                                                                                                                                                                                                                                                                                                                                                                                                               |
| status                        | enum      | ✓        |        | active, deleted                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| integration_id                | text      |          |        | Sets the **integration_id** of the term                                                                                                                                                                                                                                                                                                                                                                                                            |
| date_override_enrollment_type | text      |          |        | When set, all columns except **term_id**, **status**, **start_date**, and **end_date** will be ignored for this row. Can only be used for an existing term. If status is active, the term dates will be set to apply only to enrollments of the given type. If status is deleted, the currently set dates for the given enrollment type will be removed. Must be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, or DesignerEnrollment. |
| start_date                    | date      |          | ✓      | The date the term starts. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`                                                                                                                                                                                                                                                                                                                                                                 |
| end_date                      | date      |          | ✓      | The date the term ends. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`                                                                                                                                                                                                                                                                                                                                                                   |

## courses.csv

| Field Name             | Data Type | Required | Sticky | Description                                                                                                                                                                                                                                                             |
| ---------------------- | --------- | -------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| course_id              | text      | ✓        |        | A unique identifier used to reference courses in the enrollments data. This identifier must not change for the account, and must be globally unique. In the user interface, this is called the _SIS ID_.                                                                |
| short_name             | text      | ✓        | ✓      | A short name for the course                                                                                                                                                                                                                                             |
| long_name              | text      | ✓        | ✓      | A long name for the course. (This can be the same as the short name, but if both are available, it will provide a better user experience to provide both.)                                                                                                              |
| account_id             | text      |          | ✓      | The account identifier from accounts.csv. New courses will be attached to the root account if not specified here                                                                                                                                                        |
| term_id                | text      |          | ✓      | The term identifier from terms.csv, if no **term_id** is specified the default term for the account will be used                                                                                                                                                        |
| status                 | enum      | ✓        | ✓      | `active`, `deleted`, `completed`,`published`                                                                                                                                                                                                                            |
| integration_id         | text      |          |        | Sets the **integration_id** of the course                                                                                                                                                                                                                               |
| start_date             | date      |          | ✓      | The course start date. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`. To remove the start date pass `"<delete>"`                                                                                                                                             |
| end_date               | date      |          | ✓      | The course end date. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`. To remove the end date pass `"<delete>"`                                                                                                                                                 |
| course_format          | enum      |          |        | `on_campus`, `online`, `blended`                                                                                                                                                                                                                                        |
| blueprint_course_id    | text      |          |        | The _SIS ID_ of a pre-existing Blueprint course. When provided, the current course will be set up to receive updates from the blueprint course. Requires Blueprint Courses feature. To remove the Blueprint Course link you can pass `'dissociate'` in place of the id. |
| grade_passback_setting | text      |          | ✓      | **nightly_sync**, **not_set**                                                                                                                                                                                                                                           |
| homeroom_course        | boolean   |          |        | Whether the course is a homeroom course. Requires the courses to be associated with a "Canvas for Elementary"-enabled account.                                                                                                                                          |

If the `start_date` is set, it will override the term start date. If the
`end_date` is set, it will override the term end date.

## section.csv

| Field Name     | Data Type | Required | Sticky | Description                                                                                                                                                                                               |
| -------------- | --------- | -------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| section_id     | text      | ✓        |        | A unique identifier used to reference sections in the enrollments data. This identifier must not change for the section, and must be globally unique. In the user interface, this is called the _SIS ID_. |
| course_id      | text      | ✓        | ✓      | The course identifier from courses.csv                                                                                                                                                                    |
| name           | text      | ✓        | ✓      | The name of the section                                                                                                                                                                                   |
| status         | enum      | ✓        |        | `active`, `deleted`                                                                                                                                                                                       |
| integration_id | text      |          |        | Sets the **integration_id** of the section                                                                                                                                                                |
| start_date     | date      |          | ✓      | The section start date. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`                                                                                                                          |
| end_date       | date      |          | ✓      | The section end date The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`                                                                                                                             |

If the **start_date** is set, it will override the course and term start dates.
If the **end_date** is set, it will override the course and term end dates.

## enrollements.csv

| Field Name               | Data Type | Required | Sticky | Description                                                                                                                                                                     |
| ------------------------ | --------- | -------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| course_id                | text      | ✓*       |        | The course identifier from courses.csv                                                                                                                                          |
| root_account             | text      |          |        | The domain of the account to search for the user.                                                                                                                               |
| start_date               | date      |          | ✓      | The enrollment start date. For **start_date** to take effect the **end_date** also needs to be populated. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`              |
| end_date                 | date      |          | ✓      | The enrollment end date. For **end_date** to take effect the **start_date** also needs to be populated. The format should be in `ISO 8601: YYYY-MM-DDTHH:MM:SSZ`                |
| user_id                  | text      | ✓*       |        | The User identifier from users.csv, required to identify user. If the **user_integration_id** is present, this field will be ignored.                                           |
| user_integration_id      | text      | ✓*       |        | The **integration_id** of the user from users.csv required to identify user if the **user_id** is not present.                                                                  |
| role                     | text      | ✓*       |        | `student`, `teacher`, `ta`, `observer`, `designer`, or a `custom role` defined by the account. When using a custom role, the name is case sensitive.                            |
| role_id                  | text      | ✓*       |        | Uses a role id, either built-in or defined by the account                                                                                                                       |
| section_id               | text      | ✓*       |        | The section identifier from sections.csv, if none is specified the default section for the course will be used                                                                  |
| status                   | enum      | ✓        |        | `active`, `deleted`, `completed`, `inactive`, `deleted_last_completed`**                                                                                                        |
| associated_user_id       | text      |          |        | For observers, the user identifier from users.csv of a student in the same course that this observer should be able to see grades for. Ignored for any role other than observer |
| limit_section_privileges | boolean   |          |        | Defaults to false. When true, the enrollment will only allow the user to see and interact with users enrolled in the section given by **course_section_id**.                    |
| notify                   | boolean   |          |        | If true, a notification will be sent to the enrolled user. Notifications are not sent by default.                                                                               |

When an enrollment is in a `'completed'` state the student is limited to
read-only access to the course.

If in an `'inactive'` state, the student will be listed in the course roster for
teachers, but will not be able to view or participate in the course until the
enrollment is activated.

> **course_id** or **section_id** is required, **role** or **role_id** is
> required, and **user_id** or **user_integration_id** is required.

> **deleted_last_completed** is not a state, but it combines the deleted and
> completed states in a function that will delete an enrollment from a course if
> there are at least one other active enrollment in the course. If it is the
> last enrollment in the course it will complete it. This may be useful for when
> a user moves to a different section of a course in which there are section
> specific assignments. It offloads the logic required to determine if the
> enrollment is the users last enrollment in the given course or not.

## groups_categories.csv

| Field Name        | Data Type | Required | Sticky | Description                                                                                                                                  |
| ----------------- | --------- | -------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------- |
| group_category_id | text      | ✓        |        | A unique identifier used to reference a group category. This identifier must not change for the group category, and must be globally unique. |
| account_id        | text      |          |        | The account identifier from accounts.csv, if no account or course is specified the group will be attached to the root account.               |
| course_id         | text      |          |        | The course identifier from courses.csv, if no course or account is specified the group will be attached to the root account.                 |
| category_name     | text      | ✓        |        | The name of the group category.                                                                                                              |
| status            | enum      | ✓        |        | `active`, `deleted`                                                                                                                          |

## groups.csv

| Field Name        | Data Type | Required | Sticky | Description                                                                                                                                                                                                                     |
| ----------------- | --------- | -------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| group_id          | text      | ✓        |        | A unique identifier used to reference groups in the **group_users** data. This identifier must not change for the group, and must be globally unique.                                                                           |
| group_category_id | text      |          |        | The group category identifier from group_categories.csv, if none is specified the group will be put in the default group category for the account or course or **root_account** if there is no **course_id** or **account_id**. |
| account_id        | text      |          |        | The account identifier from accounts.csv, if none is specified the group will be attached to the root account.                                                                                                                  |
| course_id         | text      |          |        | The course identifier from courses.csv, if none is specified the group will be attached to the root account.                                                                                                                    |
| name              | text      | ✓        | ✓      | The name of the group.                                                                                                                                                                                                          |
| status            | enum      | ✓        |        | `available`, `deleted`                                                                                                                                                                                                          |

## groups_membership.csv

| Field Name | Data Type | Required | Sticky | Description                          |
| ---------- | --------- | -------- | ------ | ------------------------------------ |
| group_id   | text      | ✓        |        | The group identifier from groups.csv |
| user_id    | text      | ✓        |        | The user identifier from users.csv   |
| status     | enum      | ✓        |        | accepted, deleted                    |

## xlist.csv

| Field Name      | Data Type | Required | Sticky | Description                              |
| --------------- | --------- | -------- | ------ | ---------------------------------------- |
| xlist_course_id | text      | ✓        |        | The course identifier from courses.csv   |
| section_id      | text      | ✓        |        | The section identifier from sections.csv |
| status          | enum      | ✓        |        | active, deleted                          |

xlists.csv is optional. The goal of xlists.csv is to provide a way to add
cross-listing information to an existing course and section hierarchy. Section
ids are expected to exist already and already reference other course ids. If a
section id is provided in this file, it will be moved from its existing course
id to a new course id, such that if that new course is removed or the
cross-listing is removed, the section will revert to its previous course id. If
**xlist_course_id** does not reference an existing course, it will be created.
If you want to provide more information about the cross-listed course, please do
so in courses.csv.

While the xlists.csv does not have any sticky fields, the sections.csv does have
**course_id** as a sticky field. If the section's **course_id** is `"sticky"`,
the import will not cross list the section to another course unless it is run
with the Override UI option on the sis import.

## user_observer.csv

| Field Name  | Data Type | Required | Sticky | Description                                                |
| ----------- | --------- | -------- | ------ | ---------------------------------------------------------- |
| observer_id | text      | ✓        |        | The User identifier from users.csv for the observing user. |
| student_id  | text      | ✓        |        | The User identifier from users.csv for the student user.   |
| status      | enum      | ✓        |        | `active`, `deleted`                                        |

user_observers.csv is optional. The goal of user_observers.csv is to provide a
way to create **user_observers**. These observers will automatically be enrolled
as an observer for each of the students enrollments. When a **user_observer** is
deleted the observer enrollments of the student are also deleted.

## admins.csv

| Field Name   | Data Type | Required | Sticky | Description                                                                                                                                                             |
| ------------ | --------- | -------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| user_id      | text      | ✓        |        | The User identifier from users.csv                                                                                                                                      |
| account_id   | text      | ✓        |        | The account identifier from accounts.csv. Uses the root_account if left blank. The collumn is required even when importing for the root_account and the value is blank. |
| role_id      | text      | ✓*       |        | Uses a role id, either built-in or defined by the account.                                                                                                              |
| role         | text      | ✓*       |        | AccountAdmin, or a custom role defined by the account. When using a custom role, the name is case sensitive.                                                            |
| status       | enum      | ✓        |        | active, deleted                                                                                                                                                         |
| root_account | text      |          |        | The domain of the account to search for the user.                                                                                                                       |

admins.csv is optional. When importing admins that already exist in canvas the
admin will become managed by sis. An admin cannot be deleted by running a sis
import unless the admin is already managed by sis. Batch mode does not apply to
the admins.csv, but diffing mode does apply to the admins.csv. Admins that
already exist in the account will receive a notification of the new admin if
notification preferences are set to receive this type of notification.

> **role** or **role_id** is required.

## change_sis_id.csv

| Field Name         | Data Type | Required | Sticky | Description                                                                                                                                                                                                                              |
| ------------------ | --------- | -------- | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| old_id             | text      | ✓*       |        | The current sis_id of the object that should be changed.                                                                                                                                                                                 |
| new_id             | text      | ✓*       |        | The desired sis_id of the object. This id must be currently unique to the object type and the root_account.                                                                                                                              |
| old_integration_id | text      | ✓*       |        | The current integration_id of the object that should be changed. This column is not supported for group categories.                                                                                                                      |
| new_integration_id | text      | ✓*       |        | The desired integration_id of the object. This id must be currently unique to the object type and the root_account. This column is not supported for group categories. Can pass "<delete>" to remove the integration_id from the object. |
| type               | text      | ✓        |        | account, term, course, section, group, group_category, user                                                                                                                                                                              |

change_sis_id.csv is optional. The goal of change_sis_id.csv is to provide a way
to change **sis_ids** or **integration_ids** of existing objects. If included in
a zip file this file will process first. All other files should include the new
ids.

> **old_id** or **old_integration_id** is required, **new_id** or
> **new_integration_id** is required.

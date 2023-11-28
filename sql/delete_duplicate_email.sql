Person =
| id | email            |
| -- | ---------------- |
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |

Use Testcase
Output
| id | email            |
| -- | ---------------- |
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
Expected
| id | email            |
| -- | ---------------- |
| 1  | john@example.com |
| 2  | bob@example.com  |


-- Solution: 
DELETE p1 FROM Person p1, Person p2 WHERE p1.email = p2.email AND p1.id > p2.id;
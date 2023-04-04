The bus factor is a term for risk management. It says that how many people from a project can go under the bus, before impacting the project.
Brutal, huh?

What happens when a resource is not present in a project is that it gets delayed. Because of a loss of skill, context, how-tos, tricks, etc.

Normally, a bus factor of 1 means there are only key resources who understand the project. Loss of this resource would definitely impact the project.

A bus factor greater than 1 would mean you have less dependence on specific people.

# What are the reasons for the bus factor being one?
You don't have resources.
Resource is not interested.
 1. May be s/he doesn't belive in your capability or providing quality work.
 2. May be s/he is only interested in that project. Others have no interest in that work.
 3. May be s/he is interested in rent seeking

# When do you find out the bus factor?
Assume a key resource is on leave for more than a month. Most commonly, cron jobs are scheduled at the start or end of the month. If the job fails and we don't know how to recover, then we know we have a bus factor of 1.

# How do you increase the bus factor?
  1. Pair programming 
  1. Demos
  1. On-Call rotation
  1. Knowledge session
  1. PR/code review by people who understand the context
  1. Documentation
  1. Meeting.  Documentation complements the meeting.
  1. In a project, deliberately assigning work to a new person and keeping a specific person as backup. This helps finding out what is broken, what is not working for a new person, areas to improve.

---
title: "The One plus Three Development Cycle: Six Months In"
date: 2023-07-31
author: Graham L. Brown
draft: false
---

My team has embraced the 1+3 development cycle over the past six months, and frankly I think I like it the most of any cadence I’ve tried. From my perspective in a technical leadership role, it strikes the right balance between time to dig around in your priorities to figure out what is important, and time to execute on your priorities. 

Sound interesting? Read on.

# Structure

No good system is complete without a little structure. It’s even in the name. 1+3. A **4 week** development cycle composed of **1 planning week** and **3 work weeks**. Retros, admin work,  


## Planning

The planning week exists to determine what is important, and therefore what we as a team commit to doing, and what stretch goals exist if we finish early. The presence of stretch goals is essential as it allows the team to estimate conservatively without enabling the work to fill the time allotted. 

At the end of the planning week we should know:
 1. The most important topics to work on.
 2. The scope of a successful three week iteration.
 3. Where our time can be spent if we finish early
 4. An idea of what external resources we will need.

**Remember: 3 weeks is very short.** Whatever you choose to do, strive to keep it within the cycle.

Exactly how to arrive at the commitments will vary by team, but what I have found effective is to schedule one fifty minute meeting, your Big Rock meeting, plus an async Scope and Details meeting a few days later.

### Big Rocks, Scope and Details

The first call is your “Prioritisation” or “Big Rocks” meeting.  The entire goal of this meeting is to negotiate with your product partners, your engineering team, and your management to decide at this time, what is the most important thing? What initiative must be shifted? What is your top priority? 

I like the term “big rocks” because it demonstrates that the initiatives do not need to be resolved entirely within that cycle. Instead, just like eating an elephant a bite at a time or climbing a mountain a step at a time, a large rock gets shifted an inch at a time. To me, the iterative thinking of multiple small movements combined with the cautious approach one takes to moving a large boulder well encapsulates the aim of the Big Rock or Prioritisation meeting. Plus, attending the monthly big rock call is more fun than attending the monthly priority meeting.

As you are determining priority as a team, it is important that stakeholders come prepared. If the engineering team isn’t ready to advocate for the tech debt they see, or the product partners to advocate for the feature work, or the management team to push for other initiatives, they won’t get acted upon. Or worse, the cycle will get derailed due to the poor planning when the true priorities reveal themselves part way through the cycle. 

For myself as a technical lead on the team, that means I come armed with small and medium sized projects that I see as top priorities, why they are important, and why they are urgent. It is my goal to effectively communicate this with evidence to my coworkers and convince them these are what the team should tackle this cycle, compared to other incoming or existing work. I try to emphasize the importance of the work, the urgency of the work (Eisenhower Matrix!) plus the cost of completion and the risk of not executing. These can overlap, but each is an effective frame through which to illustrate why a stakeholder should care.

Conversely, it is role of the other stakeholders to communicate why their initiatives are important and why they are more urgent than what I bring to the table.

Finally, as the coordinator and facilitator in the room, our EM eventually gets the last say. The buck stops with them as one might say.

#### Multiple commitments

Frequently we cannot just pick a single priority to execute on. Perhaps the whole team is too large to do so, it is slow exploration work, or perhaps multiple items need to be tackled as soon as possible. 

This is normal and good. 

As our team is not small (~8 developers), we may end up working to shift multiple priorities in the same cycle. Frequently it is effective to split the team in two or three to make appropriate progress on those which we deem most important. 

### Scope and Details

The second meeting is the scope and details meeting. The goal of this meeting is to determine, given these priorities, what three weeks of work look like. Given each priority and the folks assigned to that deliverable, what does a reasonable scope of work look like? Given that scope, what would ‘done’ look like? Put another way, what is the minimum viable deliverable to be successful this cycle? What do we need to do in order to get there? These questions should lead into:
 - A definition of scope:
   - What are we doing?
   - What are we not doing?
   - What are the stretch goals we could hit if we are ahead of schedule?
 - A set of acceptance criteria
   - A ‘definition of done’
   - A shared understanding of the destination/deliverable/words failing/path

At the end of the scope and details document, we should all have a shared understanding of how each priority will be shifted, what we are committing to delivering on, what are explicitly choosing to no do, and what we could do if time provides.

These conclusions create clarity for the team and stability in the planning process for product and management. 

### Other planning week activities

Every few cycles it should contain a quick retrospective look on the cycles that have come before, and examine what we like, what we don’t like, and how we can do better. 

It’s important for the team to be heard, and it’s also important to not navel gaze too tightly. It’s good to take feedback, make change, give it a time to work, and then check in again. 

### The Work Cycle

Time to execute on your planned work! You should know at a high level what you are doing, and what done looks like. Identify the individual issues and tasks that bring us closer to done and make it happen. Read, write, investigate, code, test, MR, monitor, repeat!

In the last week, look to have everything in MR by Wednesday at the latest. That leaves a two day buffer to review, test, fix get everything through before the end of the cycle. 

If it looks like you are not going to be able to conclude before the end of the cycle, work with your manager, tech lead, and team to figure out what can be cut from the scope. How can you still deliver within the work weeks? If cutting from the scope is not possible, document your efforts and context. Priorities may shift. You want to enable any developer to pick this work back up a few cycles from now.

If completing the work inside the cycle wasn’t possible, take some time in a retro to figure out why. Did you have an over-expansive scope of work? Were you waiting on external stakeholders that couldn’t deliver on time? Or did life just get in the way? Bring the incomplete work forward into the planning meetings and decide if it is still a priority to get done. If it is, bring it forward to the next cycle. If not, leave it behind.
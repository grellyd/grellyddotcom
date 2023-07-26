---
title: "The One plus Three Development Cycle: a six month retrospective"
date: 2023-07-31
author: Graham L. Brown
draft: true
---

My team has embraced the 1+3 development cycle over the past six months, and frankly I think I like it the most of anything I’ve tried. For myself in a technical leadership role at least, it strikes the right balance between time to dig around in your priorities to figure out what is important, and time to execute on your priorities. 

Sound interesting? Read on.

# Structure

No good system is complete without a little structure. It’s even in the name. 1+3. 1 planning week, followed by 3 development weeks. 

## Planning

The planning week should contain a decision on what is important, what we as a team commit to doing, and what stretch goals exist if we finish early. The presence of stretch goals is essential as it allows the team to estimate conservatively without enabling the work to fill the time allotted. 

At the end of the planning week you should know:
 1. The most important topics to work on.
 2. The scope of a successful three week iteration.
 3. An idea of what external resources you will need.

**Remember: 3 weeks is very short.** Whatever you choose to do, work to keep it within the cycle.

### Big Rocks, Scope and Details

Exactly how to arrive at the commitments will vary by team, but what I have found effective is to schedule one fifty minute meeting, your Big Rock meeting plus an async Scope and Details meeting a few days later. The entire goal of that meeting is to negotiate with your product partners, your engineering team, and you management to decide at this time, what is the most important thing? What initiative must be shifted? What is your top priority? 

I like the term “big rocks” because it demonstrates that the initiatives do not need to be resolved entirely within that cycle. Instead, just like eating an elephant a bite at a time or climbing a mountain a step at a time, a large rock gets shifted an inch at a time. To me, the iterative thinking of multiple small movements combined with the cautious approach one takes to moving a large boulder well encapsulates the aim of the Big Rock or Priority Setting Meeting. Plus attending the monthly big rock call is more fun than attending the monthly priority meeting.  

As you are determining priority as a team, it is important that stakeholders come prepared. If the engineering team isn’t ready to advocate for the tech debt they see, or the product partners to advocate for the feature work, or the management team to push for other initiatives, they won’t get acted upon. Or worse, the cycle gets derailed due to poor planning and true priorities revealing themselves part way through the cycle. 

For myself as a technical lead on the team, that means I come armed with small and medium sized projects that I see as top priorities, why they are important, and why they are urgent. It is my goal to effectively communicate this with evidence to my coworkers and convince them these are what the team should tackle this cycle, compared to other incoming or existing work. Conversely, it is role of the other stakeholders to communicate why their initiatives are important and why they are more urgent than what I bring to the table.

Of course, as the coordinator and facilitator in the room, our EM eventually gets the last say. The buck stops with them as one might say. 

#### Scope and Details

The second meeting is the scope and details meeting. As our team is not small (~8 developers), we may end up working to shift multiple priorities in the same cycle. Frequently it is effective to split the team in two or three to make appropriate progress on that which we deem most important. 

### Other planning week activities

Every few cycles it should contain a quick retrospective look on the cycles that have come before, and examine what we like, what we don’t like, and how we can do better. 

It’s important for the team to be heard, and it’s also important to not navel gaze too tightly. It’s good to take feedback, make change, give it a time to work, and then check in again. 

How you wish to structure them in detail 

The planning week itself can be split into a couple small co

Welcome to the 3+1 development cycle!

This is a quick primer to set you expectations for the month ahead.

## Who

The ULG team. This development cycle change is limited to just the ULG team at this time. If this goes well, it could easily be expanded to other development and tech org teams.

## What

A **4 week** development cycle composed of **one planning and retro week** and **three work weeks**. 

## When

Now! The first planning week starts January 3rd, with the first full cycle starting January 6th completing January 27th. This is directly followed by the next planning and retro week (Jan 30th - Feb 3rd) for the second work period (Feb 6th - Feb 24th).

| Week | Focus |
------ | ----- |
| Jan 3rd | Planning |
| Jan 9th | Work |
| Jan 16th | Work |
| Jan 23rd | Work | 
| Jan 30th | Planning |
| Feb 6th | Work |
| ... | ... |

## How

### Planning week

At the end of the planning week you should know:
 1. The most important topics to work on.
 2. The scope of a successful three week iteration.
 3. An idea of what external resources you will need.

This week is also the time to run a retro on the previous cycle if needed.

**Remember: 3 weeks is very short.** Whatever you choose to do, work to keep it within the cycle.

### Example: The Provisioning Team

Something that has worked for the Provisioning team in the past is two meetings, each up to an hour long:

 1. Objectives & Priorities (Tuesday)
    - This goal of the meeting is to find our main focus(es) for the upcoming cycle. 
    - Without going into specific details, what matters the most right now? Are our priorities from last cycle still our top priorities for this cycle?
    - This can be thought of a 'big rocks' we need to focus on. Like with chipping away at a large stone, we do not need to resolve these topics in a cycle; sometimes it is sufficient to just move them meaningfully along.
    - Managers and folks with input into the roadmap should be present here.
 2. Scope & Detail (Thursday)
    - This meeting focuses on taking our main focuses from the previous chat and figuring out what three weeks of work on them looks like.
    - What is a reasonable scope of work for each high level priority? What does 'done' look like at the end of the cycle?
    - What is the minimum viable deliverable to be successful this cycle? How can we get there?
    - The acceptance criteria defined here will map directly to acceptance criteria in issues and subsequent MRs.
    - What are we going to do? What are we explicitly not doing? What are some stretch goals?

### Work weeks

Execute on your planned work! You should know at a high level what you are doing, and what done looks like. Identify the individual issues and tasks that bring us closer to done and make it happen. Read, write, investigate, code, test, MR, monitor, repeat!

In the last week, look to have everything in MR by Wednesday at the latest. That leaves a two day buffer to review, test, fix get everything through before the end of the cycle. 

If it looks like you are not going to be able to conclude before the end of the cycle, work with your manager, tech lead, and team to figure out what can be cut from the scope. How can you still deliver within the work weeks? If cutting from the scope is not possible, document your efforts and context. Priorities may shift. You want to enable any developer to pick this work back up a few cycles from now.

If completing the work inside the cycle wasn’t possible, take some time in a retro to figure out why. Did you have an over-expansive scope of work? Were you waiting on external stakeholders that couldn't deliver on time? Or did life just get in the way? Bring the incomplete work forward into the planning meetings and decide if it is still a priority to get done. If it is, bring it forward to the next cycle. If not, leave it behind.

## Why

Pace! The ULG dream enables us to borrow who and what we need from the different teams under our umbrella. This only becomes possible if there are natural points at which folks can move between projects.

The goal here is **not** to enforce working method compliance. You and your comrades should still pick, scope, design, and execute on the work as before. The goal is to align working periods.

# Questions?

That makes sense! Bring them up in #team-ulg and let's figure them out together. Your manager is a great resource if you need to chat privately. 

With the holidays this first cycle will probably be a bit of a mess, but that is okay. Progress, not perfection.

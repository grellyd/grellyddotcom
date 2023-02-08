---
title: "1Password SCIM Bridge Explained: what it is, and why we made it"
date: 2022-10-03
author: Graham L. Brown
draft: false
---



<aside>
⚠️This is the republication of a piece I wrote for the <a href="https://blog.1password.com/1password-scim-bridge-explained/">1Password corporate blog</a>.

I'm pretty proud of this piece, but given the fickle nature of hyperlinks, I am reproducing it here for posterity. I had help editing and refining this piece from our fantastic content team. The content is all me, just a little filtered.
</aside>

# 1Password SCIM bridge explained: what it is, and why we made it

The [1Password SCIM bridge](https://support.1password.com/scim/) is a powerful tool for businesses that want to use a password manager alongside an identity provider like Okta, Rippling, or Azure Active Directory. But if you haven’t used the SCIM bridge before, you might be wondering: What exactly is it? And does my company need a SCIM bridge?

Today, we’re going to dive in and answer both of these questions. But to do so, we have to explain the problem the SCIM bridge solves.

## The problem: your time is valuable
How can you effectively provision, manage, and deprovision users in 1Password if you work for a large organization?

Imagine you’re an administrator for a Fortune 500 company. You have over 100,000 users in your directory, and management is telling you that everyone needs access to 1Password.

Now, let’s be generous and assume inviting a user, confirming their account, and placing them in the right 1Password groups takes a total of 30 seconds [via 1Password.com](https://start.1password.com/signin). Congratulations! Your new job for the next three months is going to be adding people to 1Password. And as time goes on, there will be people who change their name, join the company, and leave for other opportunities – all of which will increase your work and take up more of your time.

> it’s not practical to manage a large number of users in 1Password without some sort of automated solution.

In short, it’s not practical to manage a large number of users in 1Password without some sort of automated solution. Your time is valuable, and we want to enable you, not slow you down.

## A first step: the industry standard

Thankfully, 1Password isn’t the first or only company to tackle the problem of managing users at scale. A variety of services exist to store, manage, and act upon user identities such as Okta, Azure Active Directory, and Google Workspace. Collectively these are known as [Identity Providers](https://www.okta.com/uk/identity-101/why-your-company-needs-an-identity-provider/), or IdPs. They’re incredibly useful if set up and configured correctly, allowing a single administrator to invite thousands of users to a new app with a single click.

In addition, something like a name change will be reflected automatically in all the apps that the IdP manages. That means there’s no need for any administrator intervention. These changes automatically inform other apps via an industry standard protocol called the [System for Cross-domain Identity Management](http://www.simplecloud.info/) (SCIM). This protocol allows apps like 1Password – referred to as Service Providers, or SPs – to speak the same language as the identity providers. So when the IdP says ‘add this new user with these characteristics’, 1Password knows exactly what to do.

To make this work, we needed to build something that can understand and interact with the SCIM protocol.

## The 1Password encryption model: an identity challenge

1Password is [designed with security in mind](https://1password.com/security/). One of our security beliefs is that your private encryption keys should never come anywhere close to our servers. They’re generated and live on devices you hold and control, and never enter our possession. If we were to be hacked (which has never happened), receive government orders, or just decide to be malicious, we couldn’t gain access to your 1Password account.

This means all of your 1Password data is encrypted with a key only you possess. It’s generated on your device using your email address, account password, and Secret Key.

But this security creates an added challenge: if your personal encryption key is stored on your device, how can 1Password and IdPs automatically carry out SCIM-related operations? After all, you use your encryption keys on your local device every time you access your account, which then allows you to invite team members, modify group memberships, and remove users.

> If your personal encryption key is stored on your device, how can 1Password and IdPs automatically carry out SCIM-related operations?

To handle requests from your identity provider, another encryption key – stored securely in a location you control – was required to access the encrypted data on our servers. Another problem was that identity providers can’t speak using encryption keys. How could we convert SCIM commands to encryption key-based operations?

## Enter stage left: the SCIM bridge

The SCIM bridge solves these problems via a server that is deployed in your company’s infrastructure. This server holds one set of encryption keys and acts as a ‘bridge’ between the IdP and 1Password, converting requests from SCIM language to 1Password’s encryption key-oriented language.

> Deploying the SCIM bridge this way lets us add additional security measures such as Secure Remote Password (SRP) to the communication between the bridge and our servers.

Here’s an example of how the SCIM bridge works when you add someone to a group in 1Password:

1. You add a user to a group in your identity provider.
1. The identity provider sends a request to the SCIM bridge explaining that the user should be added to the specified group in 1Password.
1. The SCIM bridge reads the request, then fetches the user, group, and encrypted information it needs from 1Password.
1. The SCIM bridge uses encryption keys held on your company’s server to add the user to the group.
1. The SCIM bridge tells your identity provider the operation is complete.

## Practical 1Password account management at scale

Using the 1Password SCIM bridge makes it practical to manage 1Password at scale. The bridge cuts down on tedious and time-consuming tasks for administrators, making common tasks automatic. Sending invites, confirming users, managing group memberships, and deprovisioning users all become a thing of the past.

Once it’s been deployed, an administrator shouldn’t have to think about the SCIM bridge day-to-day. The changes you make in your identity provider will be reflected automatically in 1Password. Plus, with such deep ties to your existing identity system, you can replicate your internal directory structure in 1Password with the press of a button.

> 1Password has become automatically managed.

In addition, removing someone from your identity provider will trigger the SCIM bridge to do the same in 1Password. That means the team member will lose access to all the vaults and items that were accessible from their 1Password account moments after they are disabled.

All of this means you can spend more time on other projects that will help your team stay productive and secure. 1Password has become automatically managed.

## The bottom line

The 1Password SCIM bridge allows you to connect 1Password with your existing identity provider and, thanks to the SCIM protocol, automate tasks like user provisioning and deprovisioning. That means you don’t have to go through the manual process of inviting and managing users in 1Password. The bridge also offers other security benefits like maintaining ownership of your private keys, automating confirmation of validated users, and revoking a person’s access to 1Password as soon as they’re removed from your identity provider.

Ready to start? Open up your 1Password Business account and go to the Integrations page to enable provisioning.

Still have questions? That makes sense – this is a complicated topic! Start by reading our [support documentation](https://support.1password.com/scim/) and asking for help in our [forum](https://1password.community/categories/scim-bridge). If you’re still stuck, send an email to [integrations@1password.com](mailto:integrations@1password.com) and we’ll happily answer any questions you have.
import { api, type Attendee, type Invite } from "../api";

export interface Payload {
  attendees: Attendee[];
  invites: Invite[];
}

export async function load(password: string): Promise<Payload> {
  let attendees: Attendee[] = [];
  let invites: Invite[] = [];

  const resAttendees = await api.getAttendees(password);
  if (resAttendees) {
    attendees = resAttendees.attendees;
  } else {
    throw Error("could not load attendees");
  }

  const resInvites = await api.getInvites(password);
  if (resInvites) {
    invites = resInvites.invites;
  } else {
    throw Error("could not load invites");
  }

  return { attendees: attendees, invites: invites };
}

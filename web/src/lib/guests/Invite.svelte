<script lang="ts">
  import { api, type InviteResponse } from "../api";
  import Attendee from "./Attendee.svelte";

  let inviteCode: string = $state("");
  let isLoading: boolean = $state(false);

  let invite: InviteResponse | undefined = $state(undefined);
  let currentNote: string = $state("");

  $effect(() => {
    if (inviteCode.length != 6) {
      invite = undefined;
      return;
    }
    isLoading = true;
    api
      .getInvite(inviteCode)
      .then((data) => {
        invite = data;
        currentNote = invite.note;
        isLoading = false;
      })
      .catch((e) => {
        invite = undefined;
        isLoading = false;
      });
  });

  function updateAttendeeStatusFactory(i: number) {
    return (status: boolean | null) => {
      invite!.attendees[i].confirmed = status;
    };
  }

  function updateAttendeeIsChildFactory(i: number) {
    return (isChild: boolean) => {
      invite!.attendees[i].isChild = isChild;
    };
  }

  async function saveInvite() {
    // if note is different, save it
    if (currentNote !== invite?.note) {
      invite!.note = currentNote;
      await api.saveInviteNote(inviteCode, currentNote);
    }
    // for each attendee, save
    for (let i = 0; i < invite!.attendees.length; i++) {
      const attendee = invite!.attendees[i];
      await api.saveAttendee(attendee);
    }
  }
</script>

<div class="invite">
  <div class="invite-input">
    <span>Input the code received in the invite</span>
    <input type="text" placeholder="e.g. ABC123" bind:value={inviteCode} />
  </div>
  {#if invite}
    <div class="invite-content">
      {#each invite.attendees as attendee, i}
        <Attendee
          {attendee}
          updateStatus={updateAttendeeStatusFactory(i)}
          updateIsChild={updateAttendeeIsChildFactory(i)}
        />
      {/each}
    </div>
    <div class="invite-note">
      <input
        type="text"
        placeholder="Any observations or comments"
        bind:value={currentNote}
      />
    </div>

    <div class="invite-confirm">
      <input type="submit" value="Save" onclick={saveInvite} />
    </div>
  {/if}
</div>

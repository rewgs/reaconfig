# ReaConfig

ReaConfig is an application for working with Reaper configs. Still in early development.

Features/roadmap:
- IN PROGRESS: Make any one config active at a given time (by symlinking it to the appropriate location).
- TODO: Backup the currently-active config (including the stock one) if it's not already.
- TODO: Stores any number of Reaper configs -- these can either be directories, or git repos.
- TODO: Merge any two configs.

This is what a standard Reaper config folder looks like:
```
.
├── BR.ini
├── ColorThemes
│   ├── Default_6.0.ReaperThemeZip
│   ├── Default_7.0.ReaperThemeZip
│   └── nvk_THEME.ReaperThemeZip
├── Cursors
│   └── HeDa_pointinghand.cur
├── Data
│   ├── GM.reabank
│   ├── Grooves
│   │   ├── 16th Quantize.rgt
│   │   ├── ASR10 16'th Subz 2 bar.rgt
│   │   ├── ASR10 16'th triplet Subz 2 bar.rgt
│   │   ├── ASR10 32'nd Subz 2 bar.rgt
│   │   ├── ASR10 32'nd triplet Subz 2 bar.rgt
│   │   ├── ASR10 8'th Subz 2 bar.rgt
│   │   ├── ASR10 8'th triplet Subz 2 bar.rgt
│   │   ├── DX_16_ 50% swing.rgt
│   │   ├── DX_16_ 54% swing.rgt
│   │   ├── DX_16_ 58% swing.rgt
│   │   ├── DX_16_ 62% swing.rgt
│   │   ├── DX_16_ 66% swing.rgt
│   │   ├── DX_16_ 70% swing.rgt
│   │   ├── DX_32_ 50% swing.rgt
│   │   ├── DX_32_ 66% swing.rgt
│   │   ├── DX_32_ 83% swing.rgt
│   │   ├── Korg DDD-1-16 50%.rgt
│   │   ├── Korg DDD-1-16 54%.rgt
│   │   ├── Korg DDD-1-16 58%.rgt
│   │   ├── Korg DDD-1-16 63%.rgt
│   │   ├── Korg DDD-1-16 67%.rgt
│   │   ├── Korg DDD-1-16 71%.rgt
│   │   ├── Korg DDD-1-16 75%.rgt
│   │   ├── Korg DDD-1-16 79%.rgt
│   │   ├── Korg DDD-1-16 83%.rgt
│   │   ├── Korg DDD-1-16 88%.rgt
│   │   ├── Korg DDD-1-16T%.rgt
│   │   ├── Korg DDD-1-8T%.rgt
│   │   ├── Logic_16A.rgt
│   │   ├── Logic_16B.rgt
│   │   ├── Logic_16C.rgt
│   │   ├── Logic_16D.rgt
│   │   ├── Logic_16E.rgt
│   │   ├── Logic_16F.rgt
│   │   ├── Logic_8A.rgt
│   │   ├── Logic_8B.rgt
│   │   ├── Logic_8C.rgt
│   │   ├── Logic_8D.rgt
│   │   ├── Logic_8E.rgt
│   │   ├── Logic_8F.rgt
│   │   ├── MPC 16'th Triplet Subz 4 bar.rgt
│   │   ├── MPC 32'nd Subz 4 bar.rgt
│   │   ├── MPC 32'nd Triplet Subz 4 bar.rgt
│   │   ├── MPC 50% Subz 4 bar.rgt
│   │   ├── MPC 51% Subz 4 bar.rgt
│   │   ├── MPC 52% Subz 4 bar.rgt
│   │   ├── MPC 53% Subz 4 bar.rgt
│   │   ├── MPC 54% Subz 4 bar.rgt
│   │   ├── MPC 55% Subz 4 bar.rgt
│   │   ├── MPC 56% Subz 4 bar.rgt
│   │   ├── MPC 57% Subz 4 bar.rgt
│   │   ├── MPC 58% Subz 4 bar.rgt
│   │   ├── MPC 59% Subz 4 bar.rgt
│   │   ├── MPC 60% Subz 4 bar.rgt
│   │   ├── MPC 61% Subz 4 bar.rgt
│   │   ├── MPC 62% Subz 4 bar.rgt
│   │   ├── MPC 63% Subz 4 bar.rgt
│   │   ├── MPC 64% Subz 4 bar.rgt
│   │   ├── MPC 65% Subz 4 bar.rgt
│   │   ├── MPC 66% Subz 4 bar.rgt
│   │   ├── MPC 67% Subz 4 bar.rgt
│   │   ├── MPC 68% Subz 4 bar.rgt
│   │   ├── MPC 69% Subz 4 bar.rgt
│   │   ├── MPC 70% Subz 4 bar.rgt
│   │   ├── MPC 71% Subz 4 bar.rgt
│   │   ├── MPC 72% Subz 4 bar.rgt
│   │   ├── MPC 73% Subz 4 bar.rgt
│   │   ├── MPC 74% Subz 4 bar.rgt
│   │   ├── MPC 75% Subz 4 bar.rgt
│   │   ├── MPC 8'th Triplet Subz 4 bar.rgt
│   │   ├── SP1200_50%.rgt
│   │   ├── SP1200_50%_16T.rgt
│   │   ├── SP1200_50%_32_2bar.rgt
│   │   ├── SP1200_50%_8T.rgt
│   │   ├── SP1200_54%.rgt
│   │   ├── SP1200_54%_16T.rgt
│   │   ├── SP1200_54%_8T.rgt
│   │   ├── SP1200_58%.rgt
│   │   ├── SP1200_63%.rgt
│   │   ├── SP1200_67%.rgt
│   │   ├── SP1200_71%.rgt
│   │   └── energyXT_50%.rgt
│   ├── amp_models
│   │   ├── Dumble Overdrive Special - Tweed Deluxe.wav
│   │   ├── Fender Bassman - Fender Bassman.wav
│   │   ├── Fender Bassman - Tweed Champ.wav
│   │   ├── Fender Deluxe - Marshall Stock 70.wav
│   │   ├── Fender Deluxe - Tweed Champ.wav
│   │   ├── Marshall JCM800 - Marshall Stock 70.wav
│   │   ├── Marshall JCM800 - Matchless Chieftain.wav
│   │   ├── Marshall JTM45 - Blackface Deluxe.wav
│   │   ├── Marshall JTM45 - Matchless Chieftain.wav
│   │   ├── Mesa Boogie Clean - Blackface Twin.wav
│   │   ├── Mesa Boogie Clean - Matchless Chieftain.wav
│   │   └── Vox AC30 Non Top Boost - Vox AC30.wav
│   ├── ix_keymaps
│   │   └── 00 - Default Mapping.txt
│   ├── ix_scales
│   │   ├── Chromatic.txt
│   │   ├── Dorian.txt
│   │   ├── Harmonic Minor.txt
│   │   ├── Locrian.txt
│   │   ├── Lydian.txt
│   │   ├── Major.txt
│   │   ├── Melodic Minor.txt
│   │   ├── Mixolydian.txt
│   │   ├── Natural Minor.txt
│   │   ├── Pentatonic Major.txt
│   │   ├── Pentatonic Minor.txt
│   │   ├── Phrygian.txt
│   │   └── Whole Tone.txt
│   ├── ix_sequences
│   │   ├── 1 - All Notes.txt
│   │   ├── 2 - Accent on 1.txt
│   │   ├── 2 - Accent on 2.txt
│   │   ├── 3 - Accent on 1.txt
│   │   ├── 3 - Accent on 2.txt
│   │   ├── 3 - Accent on 3.txt
│   │   ├── 4 - Accent On 1.txt
│   │   ├── 4 - Accent On 2.txt
│   │   ├── 4 - Accent On 3.txt
│   │   ├── 4 - Accent On 4.txt
│   │   └── 5 - Accent on 5.txt
│   ├── joystick_midi
│   │   ├── generic.txt
│   │   ├── rockband_drums.txt
│   │   └── rockband_guitar.txt
│   ├── reaper_imgui_doc.html
│   ├── sample.reascale
│   ├── seqbaby_data
│   │   └── _Default Kit.txt
│   ├── toolbar_icons
│   │   ├── 150
│   │   │   ├── AI_Function_AutoDoppler_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Alt3_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Alt_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Custom_Alt_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Custom_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Render_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Render_Randomize_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Snap_Offset_Alt_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Snap_Offset_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Align_Snap_Offset_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Create_New_Folder_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_In_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_In_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_Out_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_Out_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_On_Off_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Pitch-Shift_Octave_MW_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Pitch-Shift_Semitone_MW_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_MP_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Rename_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Render_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Render_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Groups_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Groups_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Items_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Items_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Settings_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Split_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Left_Edge2_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Left_Edge3_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Left_Edge_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Right_Edge2_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Right_Edge3_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Right_Edge_MP_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt3_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt4_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt5_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt6_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt7_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_NVK_DF.png
│   │   │   ├── AI_Function_Subproject_Create_Update_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Subproject_Create_Update_Alt4_NVK_DF.png
│   │   │   ├── AI_Function_Subproject_Create_Update_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Add_Nr_Take_Marker_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Add_Nr_Take_Marker_MP_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Consolidate_Takes_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Consolidate_Takes_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Convert_5_1_To_Quad_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Copy_Take_Name_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Copy_Take_Name_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Duplicate_Select_Next_Take_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Duplicate_Select_Next_Take_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_SMART_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_SMART_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Remove_Channels_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Render_New_Take_Keep_Marker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Render_New_Take_Keep_Marker_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Reverse_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Reverse_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Search_Select_Takes_With_Text_String_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt3_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt4_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_NVK_DF.png
│   │   │   ├── animation_toolbar_armed.png
│   │   │   ├── animation_toolbar_highlight.png
│   │   │   ├── toolbar_add.png
│   │   │   ├── toolbar_audio_waveform.png
│   │   │   ├── toolbar_audio_waveform_delete_remove.png
│   │   │   ├── toolbar_audio_waveform_delete_silence.png
│   │   │   ├── toolbar_audio_waveform_digital_sample_rate.png
│   │   │   ├── toolbar_audio_waveform_disk_load.png
│   │   │   ├── toolbar_audio_waveform_folder.png
│   │   │   ├── toolbar_audio_waveform_metronome.png
│   │   │   ├── toolbar_audio_waveform_move_grid_quantize.png
│   │   │   ├── toolbar_audio_waveform_normalize_gain.png
│   │   │   ├── toolbar_audio_waveform_normalize_gain_common_locked.png
│   │   │   ├── toolbar_audio_waveform_primary_external_editor.png
│   │   │   ├── toolbar_audio_waveform_properties.png
│   │   │   ├── toolbar_audio_waveform_render_disk_mono.png
│   │   │   ├── toolbar_audio_waveform_render_disk_stereo.png
│   │   │   ├── toolbar_audio_waveform_render_effects_mono.png
│   │   │   ├── toolbar_audio_waveform_render_effects_stereo.png
│   │   │   ├── toolbar_audio_waveform_reverse.png
│   │   │   ├── toolbar_audio_waveform_secondary_external_editor.png
│   │   │   ├── toolbar_audio_waveform_selection.png
│   │   │   ├── toolbar_audio_waveform_selection_trim.png
│   │   │   ├── toolbar_audio_waveform_system.png
│   │   │   ├── toolbar_audio_waveform_time_selection_render.png
│   │   │   ├── toolbar_audio_waveform_time_selection_render_stereo.png
│   │   │   ├── toolbar_audio_waveform_transient_dynamic_split.png
│   │   │   ├── toolbar_audio_waveform_transient_dynamic_split_lines.png
│   │   │   ├── toolbar_audio_waveform_transient_dynamic_split_scissors.png
│   │   │   ├── toolbar_bass_clef_note.png
│   │   │   ├── toolbar_clip_properties.png
│   │   │   ├── toolbar_clipboard_copy.png
│   │   │   ├── toolbar_clipboard_cut.png
│   │   │   ├── toolbar_clipboard_paste.png
│   │   │   ├── toolbar_color_dynamic_volume_ff.png
│   │   │   ├── toolbar_color_item.png
│   │   │   ├── toolbar_color_item_selected.png
│   │   │   ├── toolbar_color_load_disk.png
│   │   │   ├── toolbar_color_midi_channel.png
│   │   │   ├── toolbar_color_none_delete_remove.png
│   │   │   ├── toolbar_color_note_pitch.png
│   │   │   ├── toolbar_color_properties.png
│   │   │   ├── toolbar_color_random_question.png
│   │   │   ├── toolbar_color_region.png
│   │   │   ├── toolbar_color_selecte_delete_remove.png
│   │   │   ├── toolbar_color_selected.png
│   │   │   ├── toolbar_color_source_input_channel.png
│   │   │   ├── toolbar_color_sws_extension.png
│   │   │   ├── toolbar_color_take_lane.png
│   │   │   ├── toolbar_color_track.png
│   │   │   ├── toolbar_cpu_offline.png
│   │   │   ├── toolbar_cpu_online.png
│   │   │   ├── toolbar_cpu_properties_performance.png
│   │   │   ├── toolbar_delete.png
│   │   │   ├── toolbar_disk_properties_resource_path.png
│   │   │   ├── toolbar_dotted_note.png
│   │   │   ├── toolbar_eighth_quaver_grid.png
│   │   │   ├── toolbar_eighth_quaver_note.png
│   │   │   ├── toolbar_env_auto_latch.png
│   │   │   ├── toolbar_env_auto_preview.png
│   │   │   ├── toolbar_env_auto_read.png
│   │   │   ├── toolbar_env_auto_touch.png
│   │   │   ├── toolbar_env_auto_trim.png
│   │   │   ├── toolbar_env_auto_write.png
│   │   │   ├── toolbar_envelope_copy.png
│   │   │   ├── toolbar_envelope_delete_remove.png
│   │   │   ├── toolbar_envelope_fade_shape_cycle.png
│   │   │   ├── toolbar_envelope_fade_shape_none_default_delete_remove.png
│   │   │   ├── toolbar_envelope_insert_four.png
│   │   │   ├── toolbar_envelope_item_selected.png
│   │   │   ├── toolbar_envelope_item_selected_replace.png
│   │   │   ├── toolbar_envelope_knob_parameter_volume.png
│   │   │   ├── toolbar_envelope_lock.png
│   │   │   ├── toolbar_envelope_mute.png
│   │   │   ├── toolbar_envelope_new.png
│   │   │   ├── toolbar_envelope_pan.png
│   │   │   ├── toolbar_envelope_pitch_note.png
│   │   │   ├── toolbar_envelope_point_delete_remove.png
│   │   │   ├── toolbar_envelope_point_insert.png
│   │   │   ├── toolbar_envelope_point_move_axis.png
│   │   │   ├── toolbar_envelope_point_move_down.png
│   │   │   ├── toolbar_envelope_point_move_left.png
│   │   │   ├── toolbar_envelope_point_move_right.png
│   │   │   ├── toolbar_envelope_point_move_up.png
│   │   │   ├── toolbar_envelope_point_new.png
│   │   │   ├── toolbar_envelope_point_time_selection_cut_scissors.png
│   │   │   ├── toolbar_envelope_point_time_selection_delete_remove.png
│   │   │   ├── toolbar_envelope_reduce_number_points_delete_remove.png
│   │   │   ├── toolbar_envelope_show.png
│   │   │   ├── toolbar_envelope_tempo_time_clock.png
│   │   │   ├── toolbar_envelope_time_selection.png
│   │   │   ├── toolbar_envelope_vol.png
│   │   │   ├── toolbar_ex_autoplay_off.png
│   │   │   ├── toolbar_ex_autoplay_on.png
│   │   │   ├── toolbar_ex_insert_open.png
│   │   │   ├── toolbar_ex_pitch_detect_off.png
│   │   │   ├── toolbar_ex_pitch_detect_on.png
│   │   │   ├── toolbar_ex_preserve_pitch_tempo_matching_off.png
│   │   │   ├── toolbar_ex_preserve_pitch_tempo_matching_on.png
│   │   │   ├── toolbar_ex_properties_for_current_media_off.png
│   │   │   ├── toolbar_ex_properties_for_current_media_on.png
│   │   │   ├── toolbar_ex_start_on_bar_off.png
│   │   │   ├── toolbar_ex_start_on_bar_on.png
│   │   │   ├── toolbar_ex_tempo_match_double_off.png
│   │   │   ├── toolbar_ex_tempo_match_double_on.png
│   │   │   ├── toolbar_ex_tempo_match_half_off.png
│   │   │   ├── toolbar_ex_tempo_match_half_on.png
│   │   │   ├── toolbar_ex_tempo_match_off.png
│   │   │   ├── toolbar_ex_tempo_match_on.png
│   │   │   ├── toolbar_folder_add_implode.png
│   │   │   ├── toolbar_folder_add_new.png
│   │   │   ├── toolbar_folder_combine.png
│   │   │   ├── toolbar_folder_delete_remove.png
│   │   │   ├── toolbar_folder_hide.png
│   │   │   ├── toolbar_folder_item_delete_remove.png
│   │   │   ├── toolbar_folder_save_disk.png
│   │   │   ├── toolbar_folder_seperate_explode.png
│   │   │   ├── toolbar_folder_show_visible.png
│   │   │   ├── toolbar_freeze_render_apply_snowflake.png
│   │   │   ├── toolbar_glue.png
│   │   │   ├── toolbar_glue_time_selection.png
│   │   │   ├── toolbar_grid_adjust_decrease.png
│   │   │   ├── toolbar_grid_adjust_increase.png
│   │   │   ├── toolbar_group_add_item.png
│   │   │   ├── toolbar_group_add_item_selected.png
│   │   │   ├── toolbar_group_explode.png
│   │   │   ├── toolbar_group_record.png
│   │   │   ├── toolbar_group_ungroup_remove_item.png
│   │   │   ├── toolbar_group_ungroup_remove_item_selected.png
│   │   │   ├── toolbar_half_minim_grid.png
│   │   │   ├── toolbar_half_minim_note.png
│   │   │   ├── toolbar_hide_mixer.png
│   │   │   ├── toolbar_hide_selected.png
│   │   │   ├── toolbar_hide_tcp.png
│   │   │   ├── toolbar_highlight_off.png
│   │   │   ├── toolbar_highlight_on.png
│   │   │   ├── toolbar_input_fx_effect.png
│   │   │   ├── toolbar_item_arpeggiate.png
│   │   │   ├── toolbar_item_duplicate_copy.png
│   │   │   ├── toolbar_item_effects_fx_delete_remove.png
│   │   │   ├── toolbar_item_effects_fx_show.png
│   │   │   ├── toolbar_item_explode_lane_take.png
│   │   │   ├── toolbar_item_free_positioning.png
│   │   │   ├── toolbar_item_green_arrow_selected.png
│   │   │   ├── toolbar_item_green_arrow_selected_replace.png
│   │   │   ├── toolbar_item_implode_lane_take.png
│   │   │   ├── toolbar_item_insert_move_space.png
│   │   │   ├── toolbar_item_left_edge_grow.png
│   │   │   ├── toolbar_item_left_edge_position.png
│   │   │   ├── toolbar_item_left_edge_shrink.png
│   │   │   ├── toolbar_item_next.png
│   │   │   ├── toolbar_item_previous.png
│   │   │   ├── toolbar_item_properties.png
│   │   │   ├── toolbar_item_red_arrow_selected.png
│   │   │   ├── toolbar_item_remove_overlap.png
│   │   │   ├── toolbar_item_right_edge_grow.png
│   │   │   ├── toolbar_item_right_edge_position.png
│   │   │   ├── toolbar_item_right_edge_shrink.png
│   │   │   ├── toolbar_item_select.png
│   │   │   ├── toolbar_item_select_all.png
│   │   │   ├── toolbar_item_select_inverse.png
│   │   │   ├── toolbar_item_selected_above_move.png
│   │   │   ├── toolbar_item_selected_align.png
│   │   │   ├── toolbar_item_selected_area_select_move.png
│   │   │   ├── toolbar_item_selected_cut_scissors.png
│   │   │   ├── toolbar_item_selected_grow_grid.png
│   │   │   ├── toolbar_item_selected_move.png
│   │   │   ├── toolbar_item_selected_move_end.png
│   │   │   ├── toolbar_item_selected_move_horizontal_position_time.png
│   │   │   ├── toolbar_item_selected_move_nudge_left.png
│   │   │   ├── toolbar_item_selected_move_nudge_left_more.png
│   │   │   ├── toolbar_item_selected_move_nudge_right.png
│   │   │   ├── toolbar_item_selected_move_nudge_right_more.png
│   │   │   ├── toolbar_item_selected_move_vertical_track.png
│   │   │   ├── toolbar_item_selected_snap.png
│   │   │   ├── toolbar_item_selected_swap.png
│   │   │   ├── toolbar_item_selected_take_delete.png
│   │   │   ├── toolbar_item_selected_take_delete_invert.png
│   │   │   ├── toolbar_item_selected_take_extract.png
│   │   │   ├── toolbar_item_selected_take_insert.png
│   │   │   ├── toolbar_item_selected_take_move_down.png
│   │   │   ├── toolbar_item_selected_take_move_top.png
│   │   │   ├── toolbar_item_selected_take_move_up.png
│   │   │   ├── toolbar_item_selection_remove_contents_move_later.png
│   │   │   ├── toolbar_item_source_preferred_position_properties.png
│   │   │   ├── toolbar_item_take.png
│   │   │   ├── toolbar_item_take_explode.png
│   │   │   ├── toolbar_item_take_selected_extract.png
│   │   │   ├── toolbar_item_take_selected_lock.png
│   │   │   ├── toolbar_jog_back_rewind_little_bit.png
│   │   │   ├── toolbar_jog_forward_little_bit.png
│   │   │   ├── toolbar_knob_parameter_learn_lock.png
│   │   │   ├── toolbar_knob_parameter_visible_show.png
│   │   │   ├── toolbar_marker_delete_remove.png
│   │   │   ├── toolbar_marker_insert_new.png
│   │   │   ├── toolbar_marker_list.png
│   │   │   ├── toolbar_marker_load_disk.png
│   │   │   ├── toolbar_marker_lock.png
│   │   │   ├── toolbar_marker_next.png
│   │   │   ├── toolbar_marker_previous.png
│   │   │   ├── toolbar_marker_properties.png
│   │   │   ├── toolbar_marker_renum.png
│   │   │   ├── toolbar_marker_time_selection_delete_remoe.png
│   │   │   ├── toolbar_marker_time_tempo_delete_remove.png
│   │   │   ├── toolbar_marker_time_tempo_insert_new.png
│   │   │   ├── toolbar_marker_time_tempo_next.png
│   │   │   ├── toolbar_marker_time_tempo_previous.png
│   │   │   ├── toolbar_marker_time_tempo_properties.png
│   │   │   ├── toolbar_marquee_cursor_selection.png
│   │   │   ├── toolbar_marquee_cursor_selection_off.png
│   │   │   ├── toolbar_marquee_cursor_selection_on.png
│   │   │   ├── toolbar_midi_all_off.png
│   │   │   ├── toolbar_midi_all_on.png
│   │   │   ├── toolbar_midi_cc_above.png
│   │   │   ├── toolbar_midi_cc_below.png
│   │   │   ├── toolbar_midi_cc_explode.png
│   │   │   ├── toolbar_midi_cc_pitch_bend.png
│   │   │   ├── toolbar_midi_cc_scale_slope.png
│   │   │   ├── toolbar_midi_cc_selected_decrease.png
│   │   │   ├── toolbar_midi_cc_selected_increase.png
│   │   │   ├── toolbar_midi_cc_set_scale_fix.png
│   │   │   ├── toolbar_midi_delete_remove_none.png
│   │   │   ├── toolbar_midi_envelope.png
│   │   │   ├── toolbar_midi_events_mode_cycle.png
│   │   │   ├── toolbar_midi_events_mode_drum_diamond.png
│   │   │   ├── toolbar_midi_events_mode_drum_triangle.png
│   │   │   ├── toolbar_midi_events_mode_normal_rectangle.png
│   │   │   ├── toolbar_midi_folder.png
│   │   │   ├── toolbar_midi_hide_unused_note_rows.png
│   │   │   ├── toolbar_midi_hide_unused_unnamed_note_rows.png
│   │   │   ├── toolbar_midi_item_selected-262.png
│   │   │   ├── toolbar_midi_item_selected.png
│   │   │   ├── toolbar_midi_lengthen_note_grid_unit.png
│   │   │   ├── toolbar_midi_lengthen_note_pixel.png
│   │   │   ├── toolbar_midi_list.png
│   │   │   ├── toolbar_midi_notes_force_snap_scale.png
│   │   │   ├── toolbar_midi_panic_all_notes_off.png
│   │   │   ├── toolbar_midi_pitch_decrease_semitone.png
│   │   │   ├── toolbar_midi_pitch_increase_octave-173.png
│   │   │   ├── toolbar_midi_pitch_increase_octave.png
│   │   │   ├── toolbar_midi_pitch_increase_semitone.png
│   │   │   ├── toolbar_midi_pitch_transpose.png
│   │   │   ├── toolbar_midi_properties.png
│   │   │   ├── toolbar_midi_render_apply_audio_waveform.png
│   │   │   ├── toolbar_midi_shorten_note_grid_unit.png
│   │   │   ├── toolbar_midi_shorten_note_pixel.png
│   │   │   ├── toolbar_midi_show_all_note_rows.png
│   │   │   ├── toolbar_midi_size_full.png
│   │   │   ├── toolbar_midi_step.png
│   │   │   ├── toolbar_midi_waveform_audio.png
│   │   │   ├── toolbar_midi_zoom.png
│   │   │   ├── toolbar_misc_anarchy.png
│   │   │   ├── toolbar_misc_anchor.png
│   │   │   ├── toolbar_misc_back_left_previous.png
│   │   │   ├── toolbar_misc_back_left_previous_more.png
│   │   │   ├── toolbar_misc_bash.png
│   │   │   ├── toolbar_misc_bomb.png
│   │   │   ├── toolbar_misc_brush_broom_clean.png
│   │   │   ├── toolbar_misc_bulb_idea.png
│   │   │   ├── toolbar_misc_calculate_numeric.png
│   │   │   ├── toolbar_misc_car.png
│   │   │   ├── toolbar_misc_coffee.png
│   │   │   ├── toolbar_misc_devil.png
│   │   │   ├── toolbar_misc_down_next.png
│   │   │   ├── toolbar_misc_down_next_more.png
│   │   │   ├── toolbar_misc_drum.png
│   │   │   ├── toolbar_misc_duck.png
│   │   │   ├── toolbar_misc_explode.png
│   │   │   ├── toolbar_misc_filter.png
│   │   │   ├── toolbar_misc_finger.png
│   │   │   ├── toolbar_misc_firewire.png
│   │   │   ├── toolbar_misc_game.png
│   │   │   ├── toolbar_misc_guitar.png
│   │   │   ├── toolbar_misc_guitar_headstock.png
│   │   │   ├── toolbar_misc_gun.png
│   │   │   ├── toolbar_misc_heart.png
│   │   │   ├── toolbar_misc_horns.png
│   │   │   ├── toolbar_misc_house_home.png
│   │   │   ├── toolbar_misc_ibeam_cursor_selection.png
│   │   │   ├── toolbar_misc_jack_input_output.png
│   │   │   ├── toolbar_misc_key_lock.png
│   │   │   ├── toolbar_misc_keyboard.png
│   │   │   ├── toolbar_misc_lips.png
│   │   │   ├── toolbar_misc_mask.png
│   │   │   ├── toolbar_misc_mic.png
│   │   │   ├── toolbar_misc_mixer_control.png
│   │   │   ├── toolbar_misc_monitor_speaker.png
│   │   │   ├── toolbar_misc_mouse.png
│   │   │   ├── toolbar_misc_mouse_left_click.png
│   │   │   ├── toolbar_misc_mouse_right_click.png
│   │   │   ├── toolbar_misc_network_stream.png
│   │   │   ├── toolbar_misc_phones.png
│   │   │   ├── toolbar_misc_pointer_cursor.png
│   │   │   ├── toolbar_misc_pointer_cursor_white.png
│   │   │   ├── toolbar_misc_question_random.png
│   │   │   ├── toolbar_misc_radioactive.png
│   │   │   ├── toolbar_misc_right_forward_next.png
│   │   │   ├── toolbar_misc_right_forward_next_more.png
│   │   │   ├── toolbar_misc_run_backward.png
│   │   │   ├── toolbar_misc_run_forward.png
│   │   │   ├── toolbar_misc_saint.png
│   │   │   ├── toolbar_misc_skull_crossbones.png
│   │   │   ├── toolbar_misc_speech_note.png
│   │   │   ├── toolbar_misc_star.png
│   │   │   ├── toolbar_misc_star_green.png
│   │   │   ├── toolbar_misc_system_reaper.png
│   │   │   ├── toolbar_misc_tape.png
│   │   │   ├── toolbar_misc_tea_mmmmm_tea.png
│   │   │   ├── toolbar_misc_thought_idea.png
│   │   │   ├── toolbar_misc_toilet.png
│   │   │   ├── toolbar_misc_trash_bin.png
│   │   │   ├── toolbar_misc_up_previous.png
│   │   │   ├── toolbar_misc_up_previous_more.png
│   │   │   ├── toolbar_misc_usb.png
│   │   │   ├── toolbar_misc_walk_backward.png
│   │   │   ├── toolbar_misc_walk_forward.png
│   │   │   ├── toolbar_mute_envelope.png
│   │   │   ├── toolbar_mute_none_unmute.png
│   │   │   ├── toolbar_mute_on.png
│   │   │   ├── toolbar_note_gliss_slide.png
│   │   │   ├── toolbar_note_tie.png
│   │   │   ├── toolbar_parameter_scrub.png
│   │   │   ├── toolbar_path_primary_disk.png
│   │   │   ├── toolbar_path_primary_secondary_both_disk.png
│   │   │   ├── toolbar_path_secondary_disk.png
│   │   │   ├── toolbar_pitch_preserve_lock.png
│   │   │   ├── toolbar_preroll_clock.png
│   │   │   ├── toolbar_preroll_clock_record.png
│   │   │   ├── toolbar_project_save_as_new_disk.png
│   │   │   ├── toolbar_project_unused_delete_remove_disk.png
│   │   │   ├── toolbar_quarter_crotchet_grid.png
│   │   │   ├── toolbar_quarter_crotchet_note.png
│   │   │   ├── toolbar_razor_off.png
│   │   │   ├── toolbar_razor_on.png
│   │   │   ├── toolbar_record.png
│   │   │   ├── toolbar_record_arm_all.png
│   │   │   ├── toolbar_record_create_item_seperate_lane.png
│   │   │   ├── toolbar_record_loop_time_selection.png
│   │   │   ├── toolbar_record_next_beat_measure.png
│   │   │   ├── toolbar_record_next_marker.png
│   │   │   ├── toolbar_record_on.png
│   │   │   ├── toolbar_record_properties.png
│   │   │   ├── toolbar_record_selected_item_auto_punch.png
│   │   │   ├── toolbar_record_split_item_new_take.png
│   │   │   ├── toolbar_record_stop_delete.png
│   │   │   ├── toolbar_record_stop_save.png
│   │   │   ├── toolbar_record_time_selection_auto_punch.png
│   │   │   ├── toolbar_record_time_selection_selected_item_auto_punch.png
│   │   │   ├── toolbar_record_trim_item_behind_tape.png
│   │   │   ├── toolbar_region_delete_remove_none.png
│   │   │   ├── toolbar_region_new.png
│   │   │   ├── toolbar_region_next.png
│   │   │   ├── toolbar_region_play_loop.png
│   │   │   ├── toolbar_region_previous.png
│   │   │   ├── toolbar_region_properties.png
│   │   │   ├── toolbar_region_time_selection.png
│   │   │   ├── toolbar_remove_scissors_selection.png
│   │   │   ├── toolbar_remove_selection_none.png
│   │   │   ├── toolbar_render_effects_midi.png
│   │   │   ├── toolbar_ripple_all_off.png
│   │   │   ├── toolbar_ripple_all_on.png
│   │   │   ├── toolbar_ripple_one_off.png
│   │   │   ├── toolbar_ripple_one_on.png
│   │   │   ├── toolbar_screenset_camera_list.png
│   │   │   ├── toolbar_screenset_camera_new.png
│   │   │   ├── toolbar_screenset_camera_next.png
│   │   │   ├── toolbar_screenset_camera_previous.png
│   │   │   ├── toolbar_screenset_camera_save_disk.png
│   │   │   ├── toolbar_selection_delete_remove.png
│   │   │   ├── toolbar_selection_inverse_delete_remove.png
│   │   │   ├── toolbar_send_hide_mute.png
│   │   │   ├── toolbar_send_show_enable.png
│   │   │   ├── toolbar_shape_bezier.png
│   │   │   ├── toolbar_shape_fast_end.png
│   │   │   ├── toolbar_shape_fast_start.png
│   │   │   ├── toolbar_shape_linear.png
│   │   │   ├── toolbar_shape_square.png
│   │   │   ├── toolbar_show selected.png
│   │   │   ├── toolbar_show.png
│   │   │   ├── toolbar_show_insert.png
│   │   │   ├── toolbar_show_parameter.png
│   │   │   ├── toolbar_show_send.png
│   │   │   ├── toolbar_shuttle_back_rewind.png
│   │   │   ├── toolbar_shuttle_forward.png
│   │   │   ├── toolbar_sixteenth_semiquaver_grid.png
│   │   │   ├── toolbar_sixteenth_semiquaver_note.png
│   │   │   ├── toolbar_snap_offset_grid_move.png
│   │   │   ├── toolbar_solo_in_front_dim.png
│   │   │   ├── toolbar_solo_none_unsolo.png
│   │   │   ├── toolbar_solo_on.png
│   │   │   ├── toolbar_split_scissors.png
│   │   │   ├── toolbar_stretch_marker_delete_remove.png
│   │   │   ├── toolbar_stretch_marker_insert_new_add.png
│   │   │   ├── toolbar_stretch_marker_locking.png
│   │   │   ├── toolbar_stretch_marker_next.png
│   │   │   ├── toolbar_stretch_marker_previous.png
│   │   │   ├── toolbar_stretch_marker_snap_grid.png
│   │   │   ├── toolbar_stretch_marker_time_selection_delete_remove.png
│   │   │   ├── toolbar_stretch_marker_time_selection_new_add.png
│   │   │   ├── toolbar_stretch_marker_tonal.png
│   │   │   ├── toolbar_sws_extension.png
│   │   │   ├── toolbar_sws_extension_properties.png
│   │   │   ├── toolbar_sync_follow_play.png
│   │   │   ├── toolbar_sync_follow_record.png
│   │   │   ├── toolbar_system_external.png
│   │   │   ├── toolbar_system_properties.png
│   │   │   ├── toolbar_system_set_save_default_disk.png
│   │   │   ├── toolbar_theme_next.png
│   │   │   ├── toolbar_theme_previous.png
│   │   │   ├── toolbar_theme_refresh.png
│   │   │   ├── toolbar_thirty_second_demisemiquaver_.png
│   │   │   ├── toolbar_thirty_second_demisemiquaver_grid.png
│   │   │   ├── toolbar_time_beats.png
│   │   │   ├── toolbar_time_clock.png
│   │   │   ├── toolbar_time_clock_properties.png
│   │   │   ├── toolbar_time_hourglass.png
│   │   │   ├── toolbar_time_hours.png
│   │   │   ├── toolbar_time_measures.png
│   │   │   ├── toolbar_time_minutes.png
│   │   │   ├── toolbar_time_sample.png
│   │   │   ├── toolbar_time_seconds.png
│   │   │   ├── toolbar_time_selection_delete_remove.png
│   │   │   ├── toolbar_time_selection_fit_item_selected.png
│   │   │   ├── toolbar_time_selection_item_cut.png
│   │   │   ├── toolbar_time_selection_item_delete_remove.png
│   │   │   ├── toolbar_time_selection_item_selected_grow_expand.png
│   │   │   ├── toolbar_time_selection_left.png
│   │   │   ├── toolbar_time_selection_loop_lock.png
│   │   │   ├── toolbar_time_selection_loop_play.png
│   │   │   ├── toolbar_time_selection_new.png
│   │   │   ├── toolbar_time_selection_play.png
│   │   │   ├── toolbar_time_selection_properties.png
│   │   │   ├── toolbar_time_selection_region.png
│   │   │   ├── toolbar_time_selection_right.png
│   │   │   ├── toolbar_time_stretch.png
│   │   │   ├── toolbar_timebase_beats_position.png
│   │   │   ├── toolbar_timebase_beats_position_length_rate.png
│   │   │   ├── toolbar_timebase_time.png
│   │   │   ├── toolbar_tool_brush_paint.png
│   │   │   ├── toolbar_tool_crop.png
│   │   │   ├── toolbar_tool_erase_delete_remove.png
│   │   │   ├── toolbar_tool_hammer.png
│   │   │   ├── toolbar_tool_knife_trim.png
│   │   │   ├── toolbar_tool_pencil_draw.png
│   │   │   ├── toolbar_tool_razor_blade.png
│   │   │   ├── toolbar_tool_scissors_cut_trim.png
│   │   │   ├── toolbar_track_next.png
│   │   │   ├── toolbar_track_previous.png
│   │   │   ├── toolbar_transport_home_end.png
│   │   │   ├── toolbar_treble_clef_note.png
│   │   │   ├── toolbar_trim_scissors_selection.png
│   │   │   ├── toolbar_triplet_note.png
│   │   │   ├── toolbar_unfreeze_render_apply_snowflake.png
│   │   │   ├── toolbar_video_item_selected.png
│   │   │   ├── toolbar_video_properties.png
│   │   │   ├── toolbar_video_screen.png
│   │   │   ├── toolbar_video_sync_start.png
│   │   │   ├── toolbar_visible_mixer.png
│   │   │   ├── toolbar_visible_tcp.png
│   │   │   ├── toolbar_whole_semibreve_grid.png
│   │   │   ├── toolbar_whole_semibreve_note.png
│   │   │   ├── toolbar_window_floating_toolbar.png
│   │   │   ├── toolbar_window_fullscreen.png
│   │   │   ├── toolbar_window_tab_background_synchronize.png
│   │   │   ├── toolbar_window_tab_clip.png
│   │   │   ├── toolbar_window_tab_clock.png
│   │   │   ├── toolbar_window_tab_delete_remove.png
│   │   │   ├── toolbar_window_tab_docker_bottom.png
│   │   │   ├── toolbar_window_tab_docker_left.png
│   │   │   ├── toolbar_window_tab_docker_right.png
│   │   │   ├── toolbar_window_tab_docker_show.png
│   │   │   ├── toolbar_window_tab_docker_top.png
│   │   │   ├── toolbar_window_tab_effects.png
│   │   │   ├── toolbar_window_tab_folder.png
│   │   │   ├── toolbar_window_tab_list.png
│   │   │   ├── toolbar_window_tab_midi_editor.png
│   │   │   ├── toolbar_window_tab_mixer_mcp.png
│   │   │   ├── toolbar_window_tab_mixer_tcp.png
│   │   │   ├── toolbar_window_tab_navigator.png
│   │   │   ├── toolbar_window_tab_new.png
│   │   │   ├── toolbar_window_tab_new_background.png
│   │   │   ├── toolbar_window_tab_next_background.png
│   │   │   ├── toolbar_window_tab_performance.png
│   │   │   ├── toolbar_window_tab_properties.png
│   │   │   ├── toolbar_window_tab_region.png
│   │   │   ├── toolbar_window_tab_routing_matrix.png
│   │   │   ├── toolbar_window_tab_screenset_layout.png
│   │   │   ├── toolbar_window_tab_undo_history.png
│   │   │   ├── toolbar_window_tab_video.png
│   │   │   ├── toolbar_zoom_all.png
│   │   │   ├── toolbar_zoom_in_audio_waveform.png
│   │   │   ├── toolbar_zoom_in_selected_item.png
│   │   │   ├── toolbar_zoom_out_all.png
│   │   │   ├── toolbar_zoom_out_audio_waveform.png
│   │   │   ├── toolbar_zoom_project.png
│   │   │   ├── toolbar_zoom_region.png
│   │   │   ├── toolbar_zoom_selected item.png
│   │   │   ├── toolbar_zoom_selected.png
│   │   │   └── toolbar_zoom_time_selection.png
│   │   ├── 200
│   │   │   ├── AI_Function_AutoDoppler_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Alt3_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Alt_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Custom_Alt_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Custom_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Render_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Quick_Render_Randomize_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Snap_Offset_Alt_NVK_DF.png
│   │   │   ├── AI_Function_AutoDoppler_Snap_Offset_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Align_Snap_Offset_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Create_New_Folder_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_In_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_In_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_Out_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Fade_Out_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_On_Off_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Pitch-Shift_Octave_MW_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Pitch-Shift_Semitone_MW_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_MP_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Remove_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Rename_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Render_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Render_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Groups_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Groups_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Items_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Reposition_Items_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Settings_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Split_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Left_Edge2_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Left_Edge3_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Left_Edge_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Right_Edge2_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Right_Edge3_MP_NVK_DF.png
│   │   │   ├── AI_Function_Folder_Items_Trim_Right_Edge_MP_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt3_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt4_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt5_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt6_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt7_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Loopmaker_NVK_DF.png
│   │   │   ├── AI_Function_Subproject_Create_Update_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Subproject_Create_Update_Alt4_NVK_DF.png
│   │   │   ├── AI_Function_Subproject_Create_Update_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Add_Nr_Take_Marker_MP_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Add_Nr_Take_Marker_MP_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Consolidate_Takes_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Consolidate_Takes_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Convert_5_1_To_Quad_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Copy_Take_Name_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Copy_Take_Name_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Duplicate_Select_Next_Take_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Duplicate_Select_Next_Take_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_SMART_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Next_Take_Marker_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_SMART_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Previous_Take_Marker_SMART_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Remove_Channels_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Render_New_Take_Keep_Marker_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Render_New_Take_Keep_Marker_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Reverse_Alt2_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Reverse_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Search_Select_Takes_With_Text_String_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt3_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt4_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt_NVK_DF.png
│   │   │   ├── AI_Function_Takes_Toogle_Width_FX_NVK_DF.png
│   │   │   ├── animation_toolbar_armed.png
│   │   │   ├── animation_toolbar_highlight.png
│   │   │   ├── toolbar_add.png
│   │   │   ├── toolbar_audio_waveform.png
│   │   │   ├── toolbar_audio_waveform_delete_remove.png
│   │   │   ├── toolbar_audio_waveform_delete_silence.png
│   │   │   ├── toolbar_audio_waveform_digital_sample_rate.png
│   │   │   ├── toolbar_audio_waveform_disk_load.png
│   │   │   ├── toolbar_audio_waveform_folder.png
│   │   │   ├── toolbar_audio_waveform_metronome.png
│   │   │   ├── toolbar_audio_waveform_move_grid_quantize.png
│   │   │   ├── toolbar_audio_waveform_normalize_gain.png
│   │   │   ├── toolbar_audio_waveform_normalize_gain_common_locked.png
│   │   │   ├── toolbar_audio_waveform_primary_external_editor.png
│   │   │   ├── toolbar_audio_waveform_properties.png
│   │   │   ├── toolbar_audio_waveform_render_disk_mono.png
│   │   │   ├── toolbar_audio_waveform_render_disk_stereo.png
│   │   │   ├── toolbar_audio_waveform_render_effects_mono.png
│   │   │   ├── toolbar_audio_waveform_render_effects_stereo.png
│   │   │   ├── toolbar_audio_waveform_reverse.png
│   │   │   ├── toolbar_audio_waveform_secondary_external_editor.png
│   │   │   ├── toolbar_audio_waveform_selection.png
│   │   │   ├── toolbar_audio_waveform_selection_trim.png
│   │   │   ├── toolbar_audio_waveform_system.png
│   │   │   ├── toolbar_audio_waveform_time_selection_render.png
│   │   │   ├── toolbar_audio_waveform_time_selection_render_stereo.png
│   │   │   ├── toolbar_audio_waveform_transient_dynamic_split.png
│   │   │   ├── toolbar_audio_waveform_transient_dynamic_split_lines.png
│   │   │   ├── toolbar_audio_waveform_transient_dynamic_split_scissors.png
│   │   │   ├── toolbar_bass_clef_note.png
│   │   │   ├── toolbar_clip_properties.png
│   │   │   ├── toolbar_clipboard_copy.png
│   │   │   ├── toolbar_clipboard_cut.png
│   │   │   ├── toolbar_clipboard_paste.png
│   │   │   ├── toolbar_color_dynamic_volume_ff.png
│   │   │   ├── toolbar_color_item.png
│   │   │   ├── toolbar_color_item_selected.png
│   │   │   ├── toolbar_color_load_disk.png
│   │   │   ├── toolbar_color_midi_channel.png
│   │   │   ├── toolbar_color_none_delete_remove.png
│   │   │   ├── toolbar_color_note_pitch.png
│   │   │   ├── toolbar_color_properties.png
│   │   │   ├── toolbar_color_random_question.png
│   │   │   ├── toolbar_color_region.png
│   │   │   ├── toolbar_color_selecte_delete_remove.png
│   │   │   ├── toolbar_color_selected.png
│   │   │   ├── toolbar_color_source_input_channel.png
│   │   │   ├── toolbar_color_sws_extension.png
│   │   │   ├── toolbar_color_take_lane.png
│   │   │   ├── toolbar_color_track.png
│   │   │   ├── toolbar_cpu_offline.png
│   │   │   ├── toolbar_cpu_online.png
│   │   │   ├── toolbar_cpu_properties_performance.png
│   │   │   ├── toolbar_delete.png
│   │   │   ├── toolbar_disk_properties_resource_path.png
│   │   │   ├── toolbar_dotted_note.png
│   │   │   ├── toolbar_eighth_quaver_grid.png
│   │   │   ├── toolbar_eighth_quaver_note.png
│   │   │   ├── toolbar_env_auto_latch.png
│   │   │   ├── toolbar_env_auto_preview.png
│   │   │   ├── toolbar_env_auto_read.png
│   │   │   ├── toolbar_env_auto_touch.png
│   │   │   ├── toolbar_env_auto_trim.png
│   │   │   ├── toolbar_env_auto_write.png
│   │   │   ├── toolbar_envelope_copy.png
│   │   │   ├── toolbar_envelope_delete_remove.png
│   │   │   ├── toolbar_envelope_fade_shape_cycle.png
│   │   │   ├── toolbar_envelope_fade_shape_none_default_delete_remove.png
│   │   │   ├── toolbar_envelope_insert_four.png
│   │   │   ├── toolbar_envelope_item_selected.png
│   │   │   ├── toolbar_envelope_item_selected_replace.png
│   │   │   ├── toolbar_envelope_knob_parameter_volume.png
│   │   │   ├── toolbar_envelope_lock.png
│   │   │   ├── toolbar_envelope_mute.png
│   │   │   ├── toolbar_envelope_new.png
│   │   │   ├── toolbar_envelope_pan.png
│   │   │   ├── toolbar_envelope_pitch_note.png
│   │   │   ├── toolbar_envelope_point_delete_remove.png
│   │   │   ├── toolbar_envelope_point_insert.png
│   │   │   ├── toolbar_envelope_point_move_axis.png
│   │   │   ├── toolbar_envelope_point_move_down.png
│   │   │   ├── toolbar_envelope_point_move_left.png
│   │   │   ├── toolbar_envelope_point_move_right.png
│   │   │   ├── toolbar_envelope_point_move_up.png
│   │   │   ├── toolbar_envelope_point_new.png
│   │   │   ├── toolbar_envelope_point_time_selection_cut_scissors.png
│   │   │   ├── toolbar_envelope_point_time_selection_delete_remove.png
│   │   │   ├── toolbar_envelope_reduce_number_points_delete_remove.png
│   │   │   ├── toolbar_envelope_show.png
│   │   │   ├── toolbar_envelope_tempo_time_clock.png
│   │   │   ├── toolbar_envelope_time_selection.png
│   │   │   ├── toolbar_envelope_vol.png
│   │   │   ├── toolbar_ex_autoplay_off.png
│   │   │   ├── toolbar_ex_autoplay_on.png
│   │   │   ├── toolbar_ex_insert_open.png
│   │   │   ├── toolbar_ex_pitch_detect_off.png
│   │   │   ├── toolbar_ex_pitch_detect_on.png
│   │   │   ├── toolbar_ex_preserve_pitch_tempo_matching_off.png
│   │   │   ├── toolbar_ex_preserve_pitch_tempo_matching_on.png
│   │   │   ├── toolbar_ex_properties_for_current_media_off.png
│   │   │   ├── toolbar_ex_properties_for_current_media_on.png
│   │   │   ├── toolbar_ex_start_on_bar_off.png
│   │   │   ├── toolbar_ex_start_on_bar_on.png
│   │   │   ├── toolbar_ex_tempo_match_double_off.png
│   │   │   ├── toolbar_ex_tempo_match_double_on.png
│   │   │   ├── toolbar_ex_tempo_match_half_off.png
│   │   │   ├── toolbar_ex_tempo_match_half_on.png
│   │   │   ├── toolbar_ex_tempo_match_off.png
│   │   │   ├── toolbar_ex_tempo_match_on.png
│   │   │   ├── toolbar_folder_add_implode.png
│   │   │   ├── toolbar_folder_add_new.png
│   │   │   ├── toolbar_folder_combine.png
│   │   │   ├── toolbar_folder_delete_remove.png
│   │   │   ├── toolbar_folder_hide.png
│   │   │   ├── toolbar_folder_item_delete_remove.png
│   │   │   ├── toolbar_folder_save_disk.png
│   │   │   ├── toolbar_folder_seperate_explode.png
│   │   │   ├── toolbar_folder_show_visible.png
│   │   │   ├── toolbar_freeze_render_apply_snowflake.png
│   │   │   ├── toolbar_glue.png
│   │   │   ├── toolbar_glue_time_selection.png
│   │   │   ├── toolbar_grid_adjust_decrease.png
│   │   │   ├── toolbar_grid_adjust_increase.png
│   │   │   ├── toolbar_group_add_item.png
│   │   │   ├── toolbar_group_add_item_selected.png
│   │   │   ├── toolbar_group_explode.png
│   │   │   ├── toolbar_group_record.png
│   │   │   ├── toolbar_group_ungroup_remove_item.png
│   │   │   ├── toolbar_group_ungroup_remove_item_selected.png
│   │   │   ├── toolbar_half_minim_grid.png
│   │   │   ├── toolbar_half_minim_note.png
│   │   │   ├── toolbar_hide_mixer.png
│   │   │   ├── toolbar_hide_selected.png
│   │   │   ├── toolbar_hide_tcp.png
│   │   │   ├── toolbar_highlight_off.png
│   │   │   ├── toolbar_highlight_on.png
│   │   │   ├── toolbar_input_fx_effect.png
│   │   │   ├── toolbar_item_arpeggiate.png
│   │   │   ├── toolbar_item_duplicate_copy.png
│   │   │   ├── toolbar_item_effects_fx_delete_remove.png
│   │   │   ├── toolbar_item_effects_fx_show.png
│   │   │   ├── toolbar_item_explode_lane_take.png
│   │   │   ├── toolbar_item_free_positioning.png
│   │   │   ├── toolbar_item_green_arrow_selected.png
│   │   │   ├── toolbar_item_green_arrow_selected_replace.png
│   │   │   ├── toolbar_item_implode_lane_take.png
│   │   │   ├── toolbar_item_insert_move_space.png
│   │   │   ├── toolbar_item_left_edge_grow.png
│   │   │   ├── toolbar_item_left_edge_position.png
│   │   │   ├── toolbar_item_left_edge_shrink.png
│   │   │   ├── toolbar_item_next.png
│   │   │   ├── toolbar_item_previous.png
│   │   │   ├── toolbar_item_properties.png
│   │   │   ├── toolbar_item_red_arrow_selected.png
│   │   │   ├── toolbar_item_remove_overlap.png
│   │   │   ├── toolbar_item_right_edge_grow.png
│   │   │   ├── toolbar_item_right_edge_position.png
│   │   │   ├── toolbar_item_right_edge_shrink.png
│   │   │   ├── toolbar_item_select.png
│   │   │   ├── toolbar_item_select_all.png
│   │   │   ├── toolbar_item_select_inverse.png
│   │   │   ├── toolbar_item_selected_above_move.png
│   │   │   ├── toolbar_item_selected_align.png
│   │   │   ├── toolbar_item_selected_area_select_move.png
│   │   │   ├── toolbar_item_selected_cut_scissors.png
│   │   │   ├── toolbar_item_selected_grow_grid.png
│   │   │   ├── toolbar_item_selected_move.png
│   │   │   ├── toolbar_item_selected_move_end.png
│   │   │   ├── toolbar_item_selected_move_horizontal_position_time.png
│   │   │   ├── toolbar_item_selected_move_nudge_left.png
│   │   │   ├── toolbar_item_selected_move_nudge_left_more.png
│   │   │   ├── toolbar_item_selected_move_nudge_right.png
│   │   │   ├── toolbar_item_selected_move_nudge_right_more.png
│   │   │   ├── toolbar_item_selected_move_vertical_track.png
│   │   │   ├── toolbar_item_selected_snap.png
│   │   │   ├── toolbar_item_selected_swap.png
│   │   │   ├── toolbar_item_selected_take_delete.png
│   │   │   ├── toolbar_item_selected_take_delete_invert.png
│   │   │   ├── toolbar_item_selected_take_extract.png
│   │   │   ├── toolbar_item_selected_take_insert.png
│   │   │   ├── toolbar_item_selected_take_move_down.png
│   │   │   ├── toolbar_item_selected_take_move_top.png
│   │   │   ├── toolbar_item_selected_take_move_up.png
│   │   │   ├── toolbar_item_selection_remove_contents_move_later.png
│   │   │   ├── toolbar_item_source_preferred_position_properties.png
│   │   │   ├── toolbar_item_take.png
│   │   │   ├── toolbar_item_take_explode.png
│   │   │   ├── toolbar_item_take_selected_extract.png
│   │   │   ├── toolbar_item_take_selected_lock.png
│   │   │   ├── toolbar_jog_back_rewind_little_bit.png
│   │   │   ├── toolbar_jog_forward_little_bit.png
│   │   │   ├── toolbar_knob_parameter_learn_lock.png
│   │   │   ├── toolbar_knob_parameter_visible_show.png
│   │   │   ├── toolbar_marker_delete_remove.png
│   │   │   ├── toolbar_marker_insert_new.png
│   │   │   ├── toolbar_marker_list.png
│   │   │   ├── toolbar_marker_load_disk.png
│   │   │   ├── toolbar_marker_lock.png
│   │   │   ├── toolbar_marker_next.png
│   │   │   ├── toolbar_marker_previous.png
│   │   │   ├── toolbar_marker_properties.png
│   │   │   ├── toolbar_marker_renum.png
│   │   │   ├── toolbar_marker_time_selection_delete_remoe.png
│   │   │   ├── toolbar_marker_time_tempo_delete_remove.png
│   │   │   ├── toolbar_marker_time_tempo_insert_new.png
│   │   │   ├── toolbar_marker_time_tempo_next.png
│   │   │   ├── toolbar_marker_time_tempo_previous.png
│   │   │   ├── toolbar_marker_time_tempo_properties.png
│   │   │   ├── toolbar_marquee_cursor_selection.png
│   │   │   ├── toolbar_marquee_cursor_selection_off.png
│   │   │   ├── toolbar_marquee_cursor_selection_on.png
│   │   │   ├── toolbar_midi_all_off.png
│   │   │   ├── toolbar_midi_all_on.png
│   │   │   ├── toolbar_midi_cc_above.png
│   │   │   ├── toolbar_midi_cc_below.png
│   │   │   ├── toolbar_midi_cc_explode.png
│   │   │   ├── toolbar_midi_cc_pitch_bend.png
│   │   │   ├── toolbar_midi_cc_scale_slope.png
│   │   │   ├── toolbar_midi_cc_selected_decrease.png
│   │   │   ├── toolbar_midi_cc_selected_increase.png
│   │   │   ├── toolbar_midi_cc_set_scale_fix.png
│   │   │   ├── toolbar_midi_delete_remove_none.png
│   │   │   ├── toolbar_midi_envelope.png
│   │   │   ├── toolbar_midi_events_mode_cycle.png
│   │   │   ├── toolbar_midi_events_mode_drum_diamond.png
│   │   │   ├── toolbar_midi_events_mode_drum_triangle.png
│   │   │   ├── toolbar_midi_events_mode_normal_rectangle.png
│   │   │   ├── toolbar_midi_folder.png
│   │   │   ├── toolbar_midi_hide_unused_note_rows.png
│   │   │   ├── toolbar_midi_hide_unused_unnamed_note_rows.png
│   │   │   ├── toolbar_midi_item_selected-262.png
│   │   │   ├── toolbar_midi_item_selected.png
│   │   │   ├── toolbar_midi_lengthen_note_grid_unit.png
│   │   │   ├── toolbar_midi_lengthen_note_pixel.png
│   │   │   ├── toolbar_midi_list.png
│   │   │   ├── toolbar_midi_notes_force_snap_scale.png
│   │   │   ├── toolbar_midi_panic_all_notes_off.png
│   │   │   ├── toolbar_midi_pitch_decrease_semitone.png
│   │   │   ├── toolbar_midi_pitch_increase_octave-173.png
│   │   │   ├── toolbar_midi_pitch_increase_octave.png
│   │   │   ├── toolbar_midi_pitch_increase_semitone.png
│   │   │   ├── toolbar_midi_pitch_transpose.png
│   │   │   ├── toolbar_midi_properties.png
│   │   │   ├── toolbar_midi_render_apply_audio_waveform.png
│   │   │   ├── toolbar_midi_shorten_note_grid_unit.png
│   │   │   ├── toolbar_midi_shorten_note_pixel.png
│   │   │   ├── toolbar_midi_show_all_note_rows.png
│   │   │   ├── toolbar_midi_size_full.png
│   │   │   ├── toolbar_midi_step.png
│   │   │   ├── toolbar_midi_waveform_audio.png
│   │   │   ├── toolbar_midi_zoom.png
│   │   │   ├── toolbar_misc_anarchy.png
│   │   │   ├── toolbar_misc_anchor.png
│   │   │   ├── toolbar_misc_back_left_previous.png
│   │   │   ├── toolbar_misc_back_left_previous_more.png
│   │   │   ├── toolbar_misc_bash.png
│   │   │   ├── toolbar_misc_bomb.png
│   │   │   ├── toolbar_misc_brush_broom_clean.png
│   │   │   ├── toolbar_misc_bulb_idea.png
│   │   │   ├── toolbar_misc_calculate_numeric.png
│   │   │   ├── toolbar_misc_car.png
│   │   │   ├── toolbar_misc_coffee.png
│   │   │   ├── toolbar_misc_devil.png
│   │   │   ├── toolbar_misc_down_next.png
│   │   │   ├── toolbar_misc_down_next_more.png
│   │   │   ├── toolbar_misc_drum.png
│   │   │   ├── toolbar_misc_duck.png
│   │   │   ├── toolbar_misc_explode.png
│   │   │   ├── toolbar_misc_filter.png
│   │   │   ├── toolbar_misc_finger.png
│   │   │   ├── toolbar_misc_firewire.png
│   │   │   ├── toolbar_misc_game.png
│   │   │   ├── toolbar_misc_guitar.png
│   │   │   ├── toolbar_misc_guitar_headstock.png
│   │   │   ├── toolbar_misc_gun.png
│   │   │   ├── toolbar_misc_heart.png
│   │   │   ├── toolbar_misc_horns.png
│   │   │   ├── toolbar_misc_house_home.png
│   │   │   ├── toolbar_misc_ibeam_cursor_selection.png
│   │   │   ├── toolbar_misc_jack_input_output.png
│   │   │   ├── toolbar_misc_key_lock.png
│   │   │   ├── toolbar_misc_keyboard.png
│   │   │   ├── toolbar_misc_lips.png
│   │   │   ├── toolbar_misc_mask.png
│   │   │   ├── toolbar_misc_mic.png
│   │   │   ├── toolbar_misc_mixer_control.png
│   │   │   ├── toolbar_misc_monitor_speaker.png
│   │   │   ├── toolbar_misc_mouse.png
│   │   │   ├── toolbar_misc_mouse_left_click.png
│   │   │   ├── toolbar_misc_mouse_right_click.png
│   │   │   ├── toolbar_misc_network_stream.png
│   │   │   ├── toolbar_misc_phones.png
│   │   │   ├── toolbar_misc_pointer_cursor.png
│   │   │   ├── toolbar_misc_pointer_cursor_white.png
│   │   │   ├── toolbar_misc_question_random.png
│   │   │   ├── toolbar_misc_radioactive.png
│   │   │   ├── toolbar_misc_right_forward_next.png
│   │   │   ├── toolbar_misc_right_forward_next_more.png
│   │   │   ├── toolbar_misc_run_backward.png
│   │   │   ├── toolbar_misc_run_forward.png
│   │   │   ├── toolbar_misc_saint.png
│   │   │   ├── toolbar_misc_skull_crossbones.png
│   │   │   ├── toolbar_misc_speech_note.png
│   │   │   ├── toolbar_misc_star.png
│   │   │   ├── toolbar_misc_star_green.png
│   │   │   ├── toolbar_misc_system_reaper.png
│   │   │   ├── toolbar_misc_tape.png
│   │   │   ├── toolbar_misc_tea_mmmmm_tea.png
│   │   │   ├── toolbar_misc_thought_idea.png
│   │   │   ├── toolbar_misc_toilet.png
│   │   │   ├── toolbar_misc_trash_bin.png
│   │   │   ├── toolbar_misc_up_previous.png
│   │   │   ├── toolbar_misc_up_previous_more.png
│   │   │   ├── toolbar_misc_usb.png
│   │   │   ├── toolbar_misc_walk_backward.png
│   │   │   ├── toolbar_misc_walk_forward.png
│   │   │   ├── toolbar_mute_envelope.png
│   │   │   ├── toolbar_mute_none_unmute.png
│   │   │   ├── toolbar_mute_on.png
│   │   │   ├── toolbar_note_gliss_slide.png
│   │   │   ├── toolbar_note_tie.png
│   │   │   ├── toolbar_parameter_scrub.png
│   │   │   ├── toolbar_path_primary_disk.png
│   │   │   ├── toolbar_path_primary_secondary_both_disk.png
│   │   │   ├── toolbar_path_secondary_disk.png
│   │   │   ├── toolbar_pitch_preserve_lock.png
│   │   │   ├── toolbar_preroll_clock.png
│   │   │   ├── toolbar_preroll_clock_record.png
│   │   │   ├── toolbar_project_save_as_new_disk.png
│   │   │   ├── toolbar_project_unused_delete_remove_disk.png
│   │   │   ├── toolbar_quarter_crotchet_grid.png
│   │   │   ├── toolbar_quarter_crotchet_note.png
│   │   │   ├── toolbar_razor_off.png
│   │   │   ├── toolbar_razor_on.png
│   │   │   ├── toolbar_record.png
│   │   │   ├── toolbar_record_arm_all.png
│   │   │   ├── toolbar_record_create_item_seperate_lane.png
│   │   │   ├── toolbar_record_loop_time_selection.png
│   │   │   ├── toolbar_record_next_beat_measure.png
│   │   │   ├── toolbar_record_next_marker.png
│   │   │   ├── toolbar_record_on.png
│   │   │   ├── toolbar_record_properties.png
│   │   │   ├── toolbar_record_selected_item_auto_punch.png
│   │   │   ├── toolbar_record_split_item_new_take.png
│   │   │   ├── toolbar_record_stop_delete.png
│   │   │   ├── toolbar_record_stop_save.png
│   │   │   ├── toolbar_record_time_selection_auto_punch.png
│   │   │   ├── toolbar_record_time_selection_selected_item_auto_punch.png
│   │   │   ├── toolbar_record_trim_item_behind_tape.png
│   │   │   ├── toolbar_region_delete_remove_none.png
│   │   │   ├── toolbar_region_new.png
│   │   │   ├── toolbar_region_next.png
│   │   │   ├── toolbar_region_play_loop.png
│   │   │   ├── toolbar_region_previous.png
│   │   │   ├── toolbar_region_properties.png
│   │   │   ├── toolbar_region_time_selection.png
│   │   │   ├── toolbar_remove_scissors_selection.png
│   │   │   ├── toolbar_remove_selection_none.png
│   │   │   ├── toolbar_render_effects_midi.png
│   │   │   ├── toolbar_ripple_all_off.png
│   │   │   ├── toolbar_ripple_all_on.png
│   │   │   ├── toolbar_ripple_one_off.png
│   │   │   ├── toolbar_ripple_one_on.png
│   │   │   ├── toolbar_screenset_camera_list.png
│   │   │   ├── toolbar_screenset_camera_new.png
│   │   │   ├── toolbar_screenset_camera_next.png
│   │   │   ├── toolbar_screenset_camera_previous.png
│   │   │   ├── toolbar_screenset_camera_save_disk.png
│   │   │   ├── toolbar_selection_delete_remove.png
│   │   │   ├── toolbar_selection_inverse_delete_remove.png
│   │   │   ├── toolbar_send_hide_mute.png
│   │   │   ├── toolbar_send_show_enable.png
│   │   │   ├── toolbar_shape_bezier.png
│   │   │   ├── toolbar_shape_fast_end.png
│   │   │   ├── toolbar_shape_fast_start.png
│   │   │   ├── toolbar_shape_linear.png
│   │   │   ├── toolbar_shape_square.png
│   │   │   ├── toolbar_show selected.png
│   │   │   ├── toolbar_show.png
│   │   │   ├── toolbar_show_insert.png
│   │   │   ├── toolbar_show_parameter.png
│   │   │   ├── toolbar_show_send.png
│   │   │   ├── toolbar_shuttle_back_rewind.png
│   │   │   ├── toolbar_shuttle_forward.png
│   │   │   ├── toolbar_sixteenth_semiquaver_grid.png
│   │   │   ├── toolbar_sixteenth_semiquaver_note.png
│   │   │   ├── toolbar_snap_offset_grid_move.png
│   │   │   ├── toolbar_solo_in_front_dim.png
│   │   │   ├── toolbar_solo_none_unsolo.png
│   │   │   ├── toolbar_solo_on.png
│   │   │   ├── toolbar_split_scissors.png
│   │   │   ├── toolbar_stretch_marker_delete_remove.png
│   │   │   ├── toolbar_stretch_marker_insert_new_add.png
│   │   │   ├── toolbar_stretch_marker_locking.png
│   │   │   ├── toolbar_stretch_marker_next.png
│   │   │   ├── toolbar_stretch_marker_previous.png
│   │   │   ├── toolbar_stretch_marker_snap_grid.png
│   │   │   ├── toolbar_stretch_marker_time_selection_delete_remove.png
│   │   │   ├── toolbar_stretch_marker_time_selection_new_add.png
│   │   │   ├── toolbar_stretch_marker_tonal.png
│   │   │   ├── toolbar_sws_extension.png
│   │   │   ├── toolbar_sws_extension_properties.png
│   │   │   ├── toolbar_sync_follow_play.png
│   │   │   ├── toolbar_sync_follow_record.png
│   │   │   ├── toolbar_system_external.png
│   │   │   ├── toolbar_system_properties.png
│   │   │   ├── toolbar_system_set_save_default_disk.png
│   │   │   ├── toolbar_theme_next.png
│   │   │   ├── toolbar_theme_previous.png
│   │   │   ├── toolbar_theme_refresh.png
│   │   │   ├── toolbar_thirty_second_demisemiquaver_.png
│   │   │   ├── toolbar_thirty_second_demisemiquaver_grid.png
│   │   │   ├── toolbar_time_beats.png
│   │   │   ├── toolbar_time_clock.png
│   │   │   ├── toolbar_time_clock_properties.png
│   │   │   ├── toolbar_time_hourglass.png
│   │   │   ├── toolbar_time_hours.png
│   │   │   ├── toolbar_time_measures.png
│   │   │   ├── toolbar_time_minutes.png
│   │   │   ├── toolbar_time_sample.png
│   │   │   ├── toolbar_time_seconds.png
│   │   │   ├── toolbar_time_selection_delete_remove.png
│   │   │   ├── toolbar_time_selection_fit_item_selected.png
│   │   │   ├── toolbar_time_selection_item_cut.png
│   │   │   ├── toolbar_time_selection_item_delete_remove.png
│   │   │   ├── toolbar_time_selection_item_selected_grow_expand.png
│   │   │   ├── toolbar_time_selection_left.png
│   │   │   ├── toolbar_time_selection_loop_lock.png
│   │   │   ├── toolbar_time_selection_loop_play.png
│   │   │   ├── toolbar_time_selection_new.png
│   │   │   ├── toolbar_time_selection_play.png
│   │   │   ├── toolbar_time_selection_properties.png
│   │   │   ├── toolbar_time_selection_region.png
│   │   │   ├── toolbar_time_selection_right.png
│   │   │   ├── toolbar_time_stretch.png
│   │   │   ├── toolbar_timebase_beats_position.png
│   │   │   ├── toolbar_timebase_beats_position_length_rate.png
│   │   │   ├── toolbar_timebase_time.png
│   │   │   ├── toolbar_tool_brush_paint.png
│   │   │   ├── toolbar_tool_crop.png
│   │   │   ├── toolbar_tool_erase_delete_remove.png
│   │   │   ├── toolbar_tool_hammer.png
│   │   │   ├── toolbar_tool_knife_trim.png
│   │   │   ├── toolbar_tool_pencil_draw.png
│   │   │   ├── toolbar_tool_razor_blade.png
│   │   │   ├── toolbar_tool_scissors_cut_trim.png
│   │   │   ├── toolbar_track_next.png
│   │   │   ├── toolbar_track_previous.png
│   │   │   ├── toolbar_transport_home_end.png
│   │   │   ├── toolbar_treble_clef_note.png
│   │   │   ├── toolbar_trim_scissors_selection.png
│   │   │   ├── toolbar_triplet_note.png
│   │   │   ├── toolbar_unfreeze_render_apply_snowflake.png
│   │   │   ├── toolbar_video_item_selected.png
│   │   │   ├── toolbar_video_properties.png
│   │   │   ├── toolbar_video_screen.png
│   │   │   ├── toolbar_video_sync_start.png
│   │   │   ├── toolbar_visible_mixer.png
│   │   │   ├── toolbar_visible_tcp.png
│   │   │   ├── toolbar_whole_semibreve_grid.png
│   │   │   ├── toolbar_whole_semibreve_note.png
│   │   │   ├── toolbar_window_floating_toolbar.png
│   │   │   ├── toolbar_window_fullscreen.png
│   │   │   ├── toolbar_window_tab_background_synchronize.png
│   │   │   ├── toolbar_window_tab_clip.png
│   │   │   ├── toolbar_window_tab_clock.png
│   │   │   ├── toolbar_window_tab_delete_remove.png
│   │   │   ├── toolbar_window_tab_docker_bottom.png
│   │   │   ├── toolbar_window_tab_docker_left.png
│   │   │   ├── toolbar_window_tab_docker_right.png
│   │   │   ├── toolbar_window_tab_docker_show.png
│   │   │   ├── toolbar_window_tab_docker_top.png
│   │   │   ├── toolbar_window_tab_effects.png
│   │   │   ├── toolbar_window_tab_folder.png
│   │   │   ├── toolbar_window_tab_list.png
│   │   │   ├── toolbar_window_tab_midi_editor.png
│   │   │   ├── toolbar_window_tab_mixer_mcp.png
│   │   │   ├── toolbar_window_tab_mixer_tcp.png
│   │   │   ├── toolbar_window_tab_navigator.png
│   │   │   ├── toolbar_window_tab_new.png
│   │   │   ├── toolbar_window_tab_new_background.png
│   │   │   ├── toolbar_window_tab_next_background.png
│   │   │   ├── toolbar_window_tab_performance.png
│   │   │   ├── toolbar_window_tab_properties.png
│   │   │   ├── toolbar_window_tab_region.png
│   │   │   ├── toolbar_window_tab_routing_matrix.png
│   │   │   ├── toolbar_window_tab_screenset_layout.png
│   │   │   ├── toolbar_window_tab_undo_history.png
│   │   │   ├── toolbar_window_tab_video.png
│   │   │   ├── toolbar_zoom_all.png
│   │   │   ├── toolbar_zoom_in_audio_waveform.png
│   │   │   ├── toolbar_zoom_in_selected_item.png
│   │   │   ├── toolbar_zoom_out_all.png
│   │   │   ├── toolbar_zoom_out_audio_waveform.png
│   │   │   ├── toolbar_zoom_project.png
│   │   │   ├── toolbar_zoom_region.png
│   │   │   ├── toolbar_zoom_selected item.png
│   │   │   ├── toolbar_zoom_selected.png
│   │   │   └── toolbar_zoom_time_selection.png
│   │   ├── AI_Function_AutoDoppler_Alt2_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Alt3_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Alt_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Quick_Custom_Alt_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Quick_Custom_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Quick_Render_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Quick_Render_Randomize_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Snap_Offset_Alt_NVK_DF.png
│   │   ├── AI_Function_AutoDoppler_Snap_Offset_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Align_Snap_Offset_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Create_New_Folder_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Fade_In_MP_Alt_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Fade_In_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Fade_Out_MP_Alt_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Fade_Out_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_On_Off_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Pitch-Shift_Octave_MW_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Pitch-Shift_Semitone_MW_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Remove_MP_Alt2_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Remove_MP_Alt_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Remove_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Remove_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Rename_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Render_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Render_SMART_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Reposition_Groups_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Reposition_Groups_SMART_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Reposition_Items_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Reposition_Items_SMART_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Settings_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Split_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Trim_Left_Edge2_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Trim_Left_Edge3_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Trim_Left_Edge_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Trim_Right_Edge2_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Trim_Right_Edge3_MP_NVK_DF.png
│   │   ├── AI_Function_Folder_Items_Trim_Right_Edge_MP_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt2_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt3_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt4_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt5_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt6_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt7_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_Alt_NVK_DF.png
│   │   ├── AI_Function_Loopmaker_NVK_DF.png
│   │   ├── AI_Function_Subproject_Create_Update_Alt2_NVK_DF.png
│   │   ├── AI_Function_Subproject_Create_Update_Alt4_NVK_DF.png
│   │   ├── AI_Function_Subproject_Create_Update_NVK_DF.png
│   │   ├── AI_Function_Takes_Add_Nr_Take_Marker_MP_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Add_Nr_Take_Marker_MP_NVK_DF.png
│   │   ├── AI_Function_Takes_Consolidate_Takes_NVK_DF.png
│   │   ├── AI_Function_Takes_Consolidate_Takes_SMART_NVK_DF.png
│   │   ├── AI_Function_Takes_Convert_5_1_To_Quad_NVK_DF.png
│   │   ├── AI_Function_Takes_Copy_Take_Name_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Copy_Take_Name_NVK_DF.png
│   │   ├── AI_Function_Takes_Duplicate_Select_Next_Take_NVK_DF.png
│   │   ├── AI_Function_Takes_Duplicate_Select_Next_Take_SMART_NVK_DF.png
│   │   ├── AI_Function_Takes_Next_Take_Marker_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Next_Take_Marker_NVK_DF.png
│   │   ├── AI_Function_Takes_Next_Take_Marker_SMART_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Next_Take_Marker_SMART_NVK_DF.png
│   │   ├── AI_Function_Takes_Previous_Take_Marker_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Previous_Take_Marker_NVK_DF.png
│   │   ├── AI_Function_Takes_Previous_Take_Marker_SMART_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Previous_Take_Marker_SMART_NVK_DF.png
│   │   ├── AI_Function_Takes_Remove_Channels_NVK_DF.png
│   │   ├── AI_Function_Takes_Render_New_Take_Keep_Marker_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Render_New_Take_Keep_Marker_NVK_DF.png
│   │   ├── AI_Function_Takes_Reverse_Alt2_NVK_DF.png
│   │   ├── AI_Function_Takes_Reverse_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Search_Select_Takes_With_Text_String_NVK_DF.png
│   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt3_NVK_DF.png
│   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt4_NVK_DF.png
│   │   ├── AI_Function_Takes_Toogle_Width_FX_Alt_NVK_DF.png
│   │   ├── AI_Function_Takes_Toogle_Width_FX_NVK_DF.png
│   │   ├── HeDa_ARM.png
│   │   ├── HeDa_HM.png
│   │   ├── HeDa_HSM.png
│   │   ├── HeDa_IT.png
│   │   ├── HeDa_MD.png
│   │   ├── HeDa_NF.png
│   │   ├── HeDa_NR.png
│   │   ├── HeDa_NR_2.png
│   │   ├── HeDa_NR_3.png
│   │   ├── HeDa_PJ.png
│   │   ├── HeDa_RBFX.png
│   │   ├── HeDa_RL.png
│   │   ├── HeDa_RMi.png
│   │   ├── HeDa_TI.png
│   │   ├── HeDa_TT.png
│   │   ├── HeDa_TXT_e.png
│   │   ├── HeDa_TXT_i.png
│   │   ├── animation_toolbar_armed.png
│   │   ├── animation_toolbar_highlight.png
│   │   ├── toolbar_add.png
│   │   ├── toolbar_audio_waveform.png
│   │   ├── toolbar_audio_waveform_delete_remove.png
│   │   ├── toolbar_audio_waveform_delete_silence.png
│   │   ├── toolbar_audio_waveform_digital_sample_rate.png
│   │   ├── toolbar_audio_waveform_disk_load.png
│   │   ├── toolbar_audio_waveform_folder.png
│   │   ├── toolbar_audio_waveform_metronome.png
│   │   ├── toolbar_audio_waveform_move_grid_quantize.png
│   │   ├── toolbar_audio_waveform_normalize_gain.png
│   │   ├── toolbar_audio_waveform_normalize_gain_common_locked.png
│   │   ├── toolbar_audio_waveform_primary_external_editor.png
│   │   ├── toolbar_audio_waveform_properties.png
│   │   ├── toolbar_audio_waveform_render_disk_mono.png
│   │   ├── toolbar_audio_waveform_render_disk_stereo.png
│   │   ├── toolbar_audio_waveform_render_effects_mono.png
│   │   ├── toolbar_audio_waveform_render_effects_stereo.png
│   │   ├── toolbar_audio_waveform_reverse.png
│   │   ├── toolbar_audio_waveform_secondary_external_editor.png
│   │   ├── toolbar_audio_waveform_selection.png
│   │   ├── toolbar_audio_waveform_selection_trim.png
│   │   ├── toolbar_audio_waveform_system.png
│   │   ├── toolbar_audio_waveform_time_selection_render.png
│   │   ├── toolbar_audio_waveform_time_selection_render_stereo.png
│   │   ├── toolbar_audio_waveform_transient_dynamic_split.png
│   │   ├── toolbar_audio_waveform_transient_dynamic_split_lines.png
│   │   ├── toolbar_audio_waveform_transient_dynamic_split_scissors.png
│   │   ├── toolbar_bass_clef_note.png
│   │   ├── toolbar_clip_properties.png
│   │   ├── toolbar_clipboard_copy.png
│   │   ├── toolbar_clipboard_cut.png
│   │   ├── toolbar_clipboard_paste.png
│   │   ├── toolbar_color_dynamic_volume_ff.png
│   │   ├── toolbar_color_item.png
│   │   ├── toolbar_color_item_selected.png
│   │   ├── toolbar_color_load_disk.png
│   │   ├── toolbar_color_midi_channel.png
│   │   ├── toolbar_color_none_delete_remove.png
│   │   ├── toolbar_color_note_pitch.png
│   │   ├── toolbar_color_properties.png
│   │   ├── toolbar_color_random_question.png
│   │   ├── toolbar_color_region.png
│   │   ├── toolbar_color_selecte_delete_remove.png
│   │   ├── toolbar_color_selected.png
│   │   ├── toolbar_color_source_input_channel.png
│   │   ├── toolbar_color_sws_extension.png
│   │   ├── toolbar_color_take_lane.png
│   │   ├── toolbar_color_track.png
│   │   ├── toolbar_cpu_offline.png
│   │   ├── toolbar_cpu_online.png
│   │   ├── toolbar_cpu_properties_performance.png
│   │   ├── toolbar_delete.png
│   │   ├── toolbar_disk_properties_resource_path.png
│   │   ├── toolbar_dotted_note.png
│   │   ├── toolbar_eighth_quaver_grid.png
│   │   ├── toolbar_eighth_quaver_note.png
│   │   ├── toolbar_env_auto_latch.png
│   │   ├── toolbar_env_auto_preview.png
│   │   ├── toolbar_env_auto_read.png
│   │   ├── toolbar_env_auto_touch.png
│   │   ├── toolbar_env_auto_trim.png
│   │   ├── toolbar_env_auto_write.png
│   │   ├── toolbar_envelope_copy.png
│   │   ├── toolbar_envelope_delete_remove.png
│   │   ├── toolbar_envelope_fade_shape_cycle.png
│   │   ├── toolbar_envelope_fade_shape_none_default_delete_remove.png
│   │   ├── toolbar_envelope_insert_four.png
│   │   ├── toolbar_envelope_item_selected.png
│   │   ├── toolbar_envelope_item_selected_replace.png
│   │   ├── toolbar_envelope_knob_parameter_volume.png
│   │   ├── toolbar_envelope_lock.png
│   │   ├── toolbar_envelope_mute.png
│   │   ├── toolbar_envelope_new.png
│   │   ├── toolbar_envelope_pan.png
│   │   ├── toolbar_envelope_pitch_note.png
│   │   ├── toolbar_envelope_point_delete_remove.png
│   │   ├── toolbar_envelope_point_insert.png
│   │   ├── toolbar_envelope_point_move_axis.png
│   │   ├── toolbar_envelope_point_move_down.png
│   │   ├── toolbar_envelope_point_move_left.png
│   │   ├── toolbar_envelope_point_move_right.png
│   │   ├── toolbar_envelope_point_move_up.png
│   │   ├── toolbar_envelope_point_new.png
│   │   ├── toolbar_envelope_point_time_selection_cut_scissors.png
│   │   ├── toolbar_envelope_point_time_selection_delete_remove.png
│   │   ├── toolbar_envelope_reduce_number_points_delete_remove.png
│   │   ├── toolbar_envelope_show.png
│   │   ├── toolbar_envelope_tempo_time_clock.png
│   │   ├── toolbar_envelope_time_selection.png
│   │   ├── toolbar_envelope_vol.png
│   │   ├── toolbar_ex_autoplay_off.png
│   │   ├── toolbar_ex_autoplay_on.png
│   │   ├── toolbar_ex_insert_open.png
│   │   ├── toolbar_ex_pitch_detect_off.png
│   │   ├── toolbar_ex_pitch_detect_on.png
│   │   ├── toolbar_ex_preserve_pitch_tempo_matching_off.png
│   │   ├── toolbar_ex_preserve_pitch_tempo_matching_on.png
│   │   ├── toolbar_ex_properties_for_current_media_off.png
│   │   ├── toolbar_ex_properties_for_current_media_on.png
│   │   ├── toolbar_ex_start_on_bar_off.png
│   │   ├── toolbar_ex_start_on_bar_on.png
│   │   ├── toolbar_ex_tempo_match_double_off.png
│   │   ├── toolbar_ex_tempo_match_double_on.png
│   │   ├── toolbar_ex_tempo_match_half_off.png
│   │   ├── toolbar_ex_tempo_match_half_on.png
│   │   ├── toolbar_ex_tempo_match_off.png
│   │   ├── toolbar_ex_tempo_match_on.png
│   │   ├── toolbar_folder_add_implode.png
│   │   ├── toolbar_folder_add_new.png
│   │   ├── toolbar_folder_combine.png
│   │   ├── toolbar_folder_delete_remove.png
│   │   ├── toolbar_folder_hide.png
│   │   ├── toolbar_folder_item_delete_remove.png
│   │   ├── toolbar_folder_save_disk.png
│   │   ├── toolbar_folder_seperate_explode.png
│   │   ├── toolbar_folder_show_visible.png
│   │   ├── toolbar_freeze_render_apply_snowflake.png
│   │   ├── toolbar_glue.png
│   │   ├── toolbar_glue_time_selection.png
│   │   ├── toolbar_grid_adjust_decrease.png
│   │   ├── toolbar_grid_adjust_increase.png
│   │   ├── toolbar_group_add_item.png
│   │   ├── toolbar_group_add_item_selected.png
│   │   ├── toolbar_group_explode.png
│   │   ├── toolbar_group_record.png
│   │   ├── toolbar_group_ungroup_remove_item.png
│   │   ├── toolbar_group_ungroup_remove_item_selected.png
│   │   ├── toolbar_half_minim_grid.png
│   │   ├── toolbar_half_minim_note.png
│   │   ├── toolbar_hide_mixer.png
│   │   ├── toolbar_hide_selected.png
│   │   ├── toolbar_hide_tcp.png
│   │   ├── toolbar_highlight_off.png
│   │   ├── toolbar_highlight_on.png
│   │   ├── toolbar_input_fx_effect.png
│   │   ├── toolbar_item_arpeggiate.png
│   │   ├── toolbar_item_duplicate_copy.png
│   │   ├── toolbar_item_effects_fx_delete_remove.png
│   │   ├── toolbar_item_effects_fx_show.png
│   │   ├── toolbar_item_explode_lane_take.png
│   │   ├── toolbar_item_free_positioning.png
│   │   ├── toolbar_item_green_arrow_selected.png
│   │   ├── toolbar_item_green_arrow_selected_replace.png
│   │   ├── toolbar_item_implode_lane_take.png
│   │   ├── toolbar_item_insert_move_space.png
│   │   ├── toolbar_item_left_edge_grow.png
│   │   ├── toolbar_item_left_edge_position.png
│   │   ├── toolbar_item_left_edge_shrink.png
│   │   ├── toolbar_item_next.png
│   │   ├── toolbar_item_previous.png
│   │   ├── toolbar_item_properties.png
│   │   ├── toolbar_item_red_arrow_selected.png
│   │   ├── toolbar_item_remove_overlap.png
│   │   ├── toolbar_item_right_edge_grow.png
│   │   ├── toolbar_item_right_edge_position.png
│   │   ├── toolbar_item_right_edge_shrink.png
│   │   ├── toolbar_item_select.png
│   │   ├── toolbar_item_select_all.png
│   │   ├── toolbar_item_select_inverse.png
│   │   ├── toolbar_item_selected_above_move.png
│   │   ├── toolbar_item_selected_align.png
│   │   ├── toolbar_item_selected_area_select_move.png
│   │   ├── toolbar_item_selected_cut_scissors.png
│   │   ├── toolbar_item_selected_grow_grid.png
│   │   ├── toolbar_item_selected_move.png
│   │   ├── toolbar_item_selected_move_end.png
│   │   ├── toolbar_item_selected_move_horizontal_position_time.png
│   │   ├── toolbar_item_selected_move_nudge_left.png
│   │   ├── toolbar_item_selected_move_nudge_left_more.png
│   │   ├── toolbar_item_selected_move_nudge_right.png
│   │   ├── toolbar_item_selected_move_nudge_right_more.png
│   │   ├── toolbar_item_selected_move_vertical_track.png
│   │   ├── toolbar_item_selected_snap.png
│   │   ├── toolbar_item_selected_swap.png
│   │   ├── toolbar_item_selected_take_delete.png
│   │   ├── toolbar_item_selected_take_delete_invert.png
│   │   ├── toolbar_item_selected_take_extract.png
│   │   ├── toolbar_item_selected_take_insert.png
│   │   ├── toolbar_item_selected_take_move_down.png
│   │   ├── toolbar_item_selected_take_move_top.png
│   │   ├── toolbar_item_selected_take_move_up.png
│   │   ├── toolbar_item_selection_remove_contents_move_later.png
│   │   ├── toolbar_item_source_preferred_position_properties.png
│   │   ├── toolbar_item_take.png
│   │   ├── toolbar_item_take_explode.png
│   │   ├── toolbar_item_take_selected_extract.png
│   │   ├── toolbar_item_take_selected_lock.png
│   │   ├── toolbar_jog_back_rewind_little_bit.png
│   │   ├── toolbar_jog_forward_little_bit.png
│   │   ├── toolbar_knob_parameter_learn_lock.png
│   │   ├── toolbar_knob_parameter_visible_show.png
│   │   ├── toolbar_marker_delete_remove.png
│   │   ├── toolbar_marker_insert_new.png
│   │   ├── toolbar_marker_list.png
│   │   ├── toolbar_marker_load_disk.png
│   │   ├── toolbar_marker_lock.png
│   │   ├── toolbar_marker_next.png
│   │   ├── toolbar_marker_previous.png
│   │   ├── toolbar_marker_properties.png
│   │   ├── toolbar_marker_renum.png
│   │   ├── toolbar_marker_time_selection_delete_remoe.png
│   │   ├── toolbar_marker_time_tempo_delete_remove.png
│   │   ├── toolbar_marker_time_tempo_insert_new.png
│   │   ├── toolbar_marker_time_tempo_next.png
│   │   ├── toolbar_marker_time_tempo_previous.png
│   │   ├── toolbar_marker_time_tempo_properties.png
│   │   ├── toolbar_marquee_cursor_selection.png
│   │   ├── toolbar_marquee_cursor_selection_off.png
│   │   ├── toolbar_marquee_cursor_selection_on.png
│   │   ├── toolbar_midi_all_off.png
│   │   ├── toolbar_midi_all_on.png
│   │   ├── toolbar_midi_cc_above.png
│   │   ├── toolbar_midi_cc_below.png
│   │   ├── toolbar_midi_cc_explode.png
│   │   ├── toolbar_midi_cc_pitch_bend.png
│   │   ├── toolbar_midi_cc_scale_slope.png
│   │   ├── toolbar_midi_cc_selected_decrease.png
│   │   ├── toolbar_midi_cc_selected_increase.png
│   │   ├── toolbar_midi_cc_set_scale_fix.png
│   │   ├── toolbar_midi_delete_remove_none.png
│   │   ├── toolbar_midi_envelope.png
│   │   ├── toolbar_midi_events_mode_cycle.png
│   │   ├── toolbar_midi_events_mode_drum_diamond.png
│   │   ├── toolbar_midi_events_mode_drum_triangle.png
│   │   ├── toolbar_midi_events_mode_normal_rectangle.png
│   │   ├── toolbar_midi_folder.png
│   │   ├── toolbar_midi_hide_unused_note_rows.png
│   │   ├── toolbar_midi_hide_unused_unnamed_note_rows.png
│   │   ├── toolbar_midi_item.png
│   │   ├── toolbar_midi_item_selected-262.png
│   │   ├── toolbar_midi_item_selected.png
│   │   ├── toolbar_midi_lengthen_note_grid_unit.png
│   │   ├── toolbar_midi_lengthen_note_pixel.png
│   │   ├── toolbar_midi_list.png
│   │   ├── toolbar_midi_mode_event_list.png
│   │   ├── toolbar_midi_mode_musical_notation.png
│   │   ├── toolbar_midi_mode_named_notes.png
│   │   ├── toolbar_midi_mode_piano_roll.png
│   │   ├── toolbar_midi_notes_force_snap_scale.png
│   │   ├── toolbar_midi_panic_all_notes_off.png
│   │   ├── toolbar_midi_pitch_decrease_octave.png
│   │   ├── toolbar_midi_pitch_decrease_semitone.png
│   │   ├── toolbar_midi_pitch_increase_octave-173.png
│   │   ├── toolbar_midi_pitch_increase_octave.png
│   │   ├── toolbar_midi_pitch_increase_semitone.png
│   │   ├── toolbar_midi_pitch_transpose.png
│   │   ├── toolbar_midi_properties.png
│   │   ├── toolbar_midi_render_apply_audio_waveform.png
│   │   ├── toolbar_midi_shorten_note_grid_unit.png
│   │   ├── toolbar_midi_shorten_note_pixel.png
│   │   ├── toolbar_midi_show_all_note_rows.png
│   │   ├── toolbar_midi_size_full.png
│   │   ├── toolbar_midi_step.png
│   │   ├── toolbar_midi_waveform_audio.png
│   │   ├── toolbar_midi_zoom.png
│   │   ├── toolbar_misc_anarchy.png
│   │   ├── toolbar_misc_anchor.png
│   │   ├── toolbar_misc_back_left_previous.png
│   │   ├── toolbar_misc_back_left_previous_more.png
│   │   ├── toolbar_misc_bash.png
│   │   ├── toolbar_misc_bomb.png
│   │   ├── toolbar_misc_brush_broom_clean.png
│   │   ├── toolbar_misc_bulb_idea.png
│   │   ├── toolbar_misc_calculate_numeric.png
│   │   ├── toolbar_misc_car.png
│   │   ├── toolbar_misc_coffee.png
│   │   ├── toolbar_misc_devil.png
│   │   ├── toolbar_misc_down_next.png
│   │   ├── toolbar_misc_down_next_more.png
│   │   ├── toolbar_misc_drum.png
│   │   ├── toolbar_misc_duck.png
│   │   ├── toolbar_misc_explode.png
│   │   ├── toolbar_misc_filter.png
│   │   ├── toolbar_misc_finger.png
│   │   ├── toolbar_misc_firewire.png
│   │   ├── toolbar_misc_game.png
│   │   ├── toolbar_misc_guitar.png
│   │   ├── toolbar_misc_guitar_headstock.png
│   │   ├── toolbar_misc_gun.png
│   │   ├── toolbar_misc_heart.png
│   │   ├── toolbar_misc_horns.png
│   │   ├── toolbar_misc_house_home.png
│   │   ├── toolbar_misc_ibeam_cursor_selection.png
│   │   ├── toolbar_misc_jack_input_output.png
│   │   ├── toolbar_misc_key_lock.png
│   │   ├── toolbar_misc_keyboard.png
│   │   ├── toolbar_misc_lips.png
│   │   ├── toolbar_misc_mask.png
│   │   ├── toolbar_misc_mic.png
│   │   ├── toolbar_misc_mixer_control.png
│   │   ├── toolbar_misc_monitor_speaker.png
│   │   ├── toolbar_misc_mouse.png
│   │   ├── toolbar_misc_mouse_left_click.png
│   │   ├── toolbar_misc_mouse_right_click.png
│   │   ├── toolbar_misc_network_stream.png
│   │   ├── toolbar_misc_phones.png
│   │   ├── toolbar_misc_pointer_cursor.png
│   │   ├── toolbar_misc_pointer_cursor_white.png
│   │   ├── toolbar_misc_question_random.png
│   │   ├── toolbar_misc_radioactive.png
│   │   ├── toolbar_misc_right_forward_next.png
│   │   ├── toolbar_misc_right_forward_next_more.png
│   │   ├── toolbar_misc_run_backward.png
│   │   ├── toolbar_misc_run_forward.png
│   │   ├── toolbar_misc_saint.png
│   │   ├── toolbar_misc_skull_crossbones.png
│   │   ├── toolbar_misc_speech_note.png
│   │   ├── toolbar_misc_star.png
│   │   ├── toolbar_misc_star_green.png
│   │   ├── toolbar_misc_system_reaper.png
│   │   ├── toolbar_misc_tape.png
│   │   ├── toolbar_misc_tea_mmmmm_tea.png
│   │   ├── toolbar_misc_thought_idea.png
│   │   ├── toolbar_misc_toilet.png
│   │   ├── toolbar_misc_trash_bin.png
│   │   ├── toolbar_misc_up_previous.png
│   │   ├── toolbar_misc_up_previous_more.png
│   │   ├── toolbar_misc_usb.png
│   │   ├── toolbar_misc_walk_backward.png
│   │   ├── toolbar_misc_walk_forward.png
│   │   ├── toolbar_mute_envelope.png
│   │   ├── toolbar_mute_none_unmute.png
│   │   ├── toolbar_mute_on.png
│   │   ├── toolbar_note_gliss_slide.png
│   │   ├── toolbar_note_tie.png
│   │   ├── toolbar_parameter_scrub.png
│   │   ├── toolbar_path_primary_disk.png
│   │   ├── toolbar_path_primary_secondary_both_disk.png
│   │   ├── toolbar_path_secondary_disk.png
│   │   ├── toolbar_pitch_preserve_lock.png
│   │   ├── toolbar_preroll_clock.png
│   │   ├── toolbar_preroll_clock_record.png
│   │   ├── toolbar_project_save_as_new_disk.png
│   │   ├── toolbar_project_unused_delete_remove_disk.png
│   │   ├── toolbar_quarter_crotchet_grid.png
│   │   ├── toolbar_quarter_crotchet_note.png
│   │   ├── toolbar_razor_off.png
│   │   ├── toolbar_razor_on.png
│   │   ├── toolbar_record.png
│   │   ├── toolbar_record_arm_all.png
│   │   ├── toolbar_record_create_item_seperate_lane.png
│   │   ├── toolbar_record_loop_time_selection.png
│   │   ├── toolbar_record_next_beat_measure.png
│   │   ├── toolbar_record_next_marker.png
│   │   ├── toolbar_record_on.png
│   │   ├── toolbar_record_properties.png
│   │   ├── toolbar_record_selected_item_auto_punch.png
│   │   ├── toolbar_record_split_item_new_take.png
│   │   ├── toolbar_record_stop_delete.png
│   │   ├── toolbar_record_stop_save.png
│   │   ├── toolbar_record_time_selection_auto_punch.png
│   │   ├── toolbar_record_time_selection_selected_item_auto_punch.png
│   │   ├── toolbar_record_trim_item_behind_tape.png
│   │   ├── toolbar_region_delete_remove_none.png
│   │   ├── toolbar_region_new.png
│   │   ├── toolbar_region_next.png
│   │   ├── toolbar_region_play_loop.png
│   │   ├── toolbar_region_previous.png
│   │   ├── toolbar_region_properties.png
│   │   ├── toolbar_region_time_selection.png
│   │   ├── toolbar_remove_scissors_selection.png
│   │   ├── toolbar_remove_selection_none.png
│   │   ├── toolbar_render_effects_midi.png
│   │   ├── toolbar_ripple_all_off.png
│   │   ├── toolbar_ripple_all_on.png
│   │   ├── toolbar_ripple_one_off.png
│   │   ├── toolbar_ripple_one_on.png
│   │   ├── toolbar_screenset_camera_list.png
│   │   ├── toolbar_screenset_camera_new.png
│   │   ├── toolbar_screenset_camera_next.png
│   │   ├── toolbar_screenset_camera_previous.png
│   │   ├── toolbar_screenset_camera_save_disk.png
│   │   ├── toolbar_selection_delete_remove.png
│   │   ├── toolbar_selection_inverse_delete_remove.png
│   │   ├── toolbar_send_hide_mute.png
│   │   ├── toolbar_send_show_enable.png
│   │   ├── toolbar_shape_bezier.png
│   │   ├── toolbar_shape_fast_end.png
│   │   ├── toolbar_shape_fast_start.png
│   │   ├── toolbar_shape_linear.png
│   │   ├── toolbar_shape_square.png
│   │   ├── toolbar_show selected.png
│   │   ├── toolbar_show.png
│   │   ├── toolbar_show_insert.png
│   │   ├── toolbar_show_parameter.png
│   │   ├── toolbar_show_send.png
│   │   ├── toolbar_shuttle_back_rewind.png
│   │   ├── toolbar_shuttle_forward.png
│   │   ├── toolbar_sixteenth_semiquaver_grid.png
│   │   ├── toolbar_sixteenth_semiquaver_note.png
│   │   ├── toolbar_snap_offset_grid_move.png
│   │   ├── toolbar_solo_in_front_dim.png
│   │   ├── toolbar_solo_none_unsolo.png
│   │   ├── toolbar_solo_on.png
│   │   ├── toolbar_split_scissors.png
│   │   ├── toolbar_stretch_marker_delete_remove.png
│   │   ├── toolbar_stretch_marker_insert_new_add.png
│   │   ├── toolbar_stretch_marker_locking.png
│   │   ├── toolbar_stretch_marker_next.png
│   │   ├── toolbar_stretch_marker_previous.png
│   │   ├── toolbar_stretch_marker_snap_grid.png
│   │   ├── toolbar_stretch_marker_time_selection_delete_remove.png
│   │   ├── toolbar_stretch_marker_time_selection_new_add.png
│   │   ├── toolbar_stretch_marker_tonal.png
│   │   ├── toolbar_sws_extension.png
│   │   ├── toolbar_sws_extension_properties.png
│   │   ├── toolbar_sync_follow_play.png
│   │   ├── toolbar_sync_follow_record.png
│   │   ├── toolbar_system_external.png
│   │   ├── toolbar_system_properties.png
│   │   ├── toolbar_system_set_save_default_disk.png
│   │   ├── toolbar_theme_next.png
│   │   ├── toolbar_theme_previous.png
│   │   ├── toolbar_theme_refresh.png
│   │   ├── toolbar_thirty_second_demisemiquaver_.png
│   │   ├── toolbar_thirty_second_demisemiquaver_grid.png
│   │   ├── toolbar_time_beats.png
│   │   ├── toolbar_time_clock.png
│   │   ├── toolbar_time_clock_properties.png
│   │   ├── toolbar_time_hourglass.png
│   │   ├── toolbar_time_hours.png
│   │   ├── toolbar_time_measures.png
│   │   ├── toolbar_time_minutes.png
│   │   ├── toolbar_time_sample.png
│   │   ├── toolbar_time_seconds.png
│   │   ├── toolbar_time_selection_delete_remove.png
│   │   ├── toolbar_time_selection_fit_item_selected.png
│   │   ├── toolbar_time_selection_item_cut.png
│   │   ├── toolbar_time_selection_item_delete_remove.png
│   │   ├── toolbar_time_selection_item_selected_grow_expand.png
│   │   ├── toolbar_time_selection_left.png
│   │   ├── toolbar_time_selection_loop_lock.png
│   │   ├── toolbar_time_selection_loop_play.png
│   │   ├── toolbar_time_selection_new.png
│   │   ├── toolbar_time_selection_play.png
│   │   ├── toolbar_time_selection_properties.png
│   │   ├── toolbar_time_selection_region.png
│   │   ├── toolbar_time_selection_right.png
│   │   ├── toolbar_time_stretch.png
│   │   ├── toolbar_timebase_beats_position.png
│   │   ├── toolbar_timebase_beats_position_length_rate.png
│   │   ├── toolbar_timebase_time.png
│   │   ├── toolbar_tool_brush_paint.png
│   │   ├── toolbar_tool_crop.png
│   │   ├── toolbar_tool_erase_delete_remove.png
│   │   ├── toolbar_tool_hammer.png
│   │   ├── toolbar_tool_knife_trim.png
│   │   ├── toolbar_tool_pencil_draw.png
│   │   ├── toolbar_tool_razor_blade.png
│   │   ├── toolbar_tool_scissors_cut_trim.png
│   │   ├── toolbar_track_next.png
│   │   ├── toolbar_track_previous.png
│   │   ├── toolbar_transport_home_end.png
│   │   ├── toolbar_treble_clef_note.png
│   │   ├── toolbar_trim_scissors_selection.png
│   │   ├── toolbar_triplet_note.png
│   │   ├── toolbar_unfreeze_render_apply_snowflake.png
│   │   ├── toolbar_v3_envitem_off.png
│   │   ├── toolbar_v3_envitem_on.png
│   │   ├── toolbar_v3_grid_off.png
│   │   ├── toolbar_v3_grid_on.png
│   │   ├── toolbar_v3_group_off.png
│   │   ├── toolbar_v3_group_on.png
│   │   ├── toolbar_v3_load.png
│   │   ├── toolbar_v3_lock_off.png
│   │   ├── toolbar_v3_lock_on.png
│   │   ├── toolbar_v3_metro_off.png
│   │   ├── toolbar_v3_metro_on.png
│   │   ├── toolbar_v3_new.png
│   │   ├── toolbar_v3_projprop.png
│   │   ├── toolbar_v3_redo.png
│   │   ├── toolbar_v3_ripple_all.png
│   │   ├── toolbar_v3_ripple_off.png
│   │   ├── toolbar_v3_ripple_one.png
│   │   ├── toolbar_v3_save.png
│   │   ├── toolbar_v3_snap_off.png
│   │   ├── toolbar_v3_snap_on.png
│   │   ├── toolbar_v3_undo.png
│   │   ├── toolbar_v3_xfade_off.png
│   │   ├── toolbar_v3_xfade_on.png
│   │   ├── toolbar_video_item_selected.png
│   │   ├── toolbar_video_properties.png
│   │   ├── toolbar_video_screen.png
│   │   ├── toolbar_video_sync_start.png
│   │   ├── toolbar_visible_mixer.png
│   │   ├── toolbar_visible_tcp.png
│   │   ├── toolbar_whole_semibreve_grid.png
│   │   ├── toolbar_whole_semibreve_note.png
│   │   ├── toolbar_window_floating_toolbar.png
│   │   ├── toolbar_window_fullscreen.png
│   │   ├── toolbar_window_tab_background_synchronize.png
│   │   ├── toolbar_window_tab_clip.png
│   │   ├── toolbar_window_tab_clock.png
│   │   ├── toolbar_window_tab_delete_remove.png
│   │   ├── toolbar_window_tab_docker_bottom.png
│   │   ├── toolbar_window_tab_docker_left.png
│   │   ├── toolbar_window_tab_docker_right.png
│   │   ├── toolbar_window_tab_docker_show.png
│   │   ├── toolbar_window_tab_docker_top.png
│   │   ├── toolbar_window_tab_effects.png
│   │   ├── toolbar_window_tab_folder.png
│   │   ├── toolbar_window_tab_list.png
│   │   ├── toolbar_window_tab_midi_editor.png
│   │   ├── toolbar_window_tab_mixer_mcp.png
│   │   ├── toolbar_window_tab_mixer_tcp.png
│   │   ├── toolbar_window_tab_navigator.png
│   │   ├── toolbar_window_tab_new.png
│   │   ├── toolbar_window_tab_new_background.png
│   │   ├── toolbar_window_tab_next_background.png
│   │   ├── toolbar_window_tab_performance.png
│   │   ├── toolbar_window_tab_properties.png
│   │   ├── toolbar_window_tab_region.png
│   │   ├── toolbar_window_tab_routing_matrix.png
│   │   ├── toolbar_window_tab_screenset_layout.png
│   │   ├── toolbar_window_tab_undo_history.png
│   │   ├── toolbar_window_tab_video.png
│   │   ├── toolbar_zoom_all.png
│   │   ├── toolbar_zoom_in_audio_waveform.png
│   │   ├── toolbar_zoom_in_selected_item.png
│   │   ├── toolbar_zoom_out_all.png
│   │   ├── toolbar_zoom_out_audio_waveform.png
│   │   ├── toolbar_zoom_project.png
│   │   ├── toolbar_zoom_region.png
│   │   ├── toolbar_zoom_selected item.png
│   │   ├── toolbar_zoom_selected.png
│   │   └── toolbar_zoom_time_selection.png
│   └── track_icons
│       ├── ac_guitar.png
│       ├── ac_guitar_full.png
│       ├── amp.png
│       ├── amp_combo.png
│       ├── balalaika.png
│       ├── banjo.png
│       ├── bass.png
│       ├── bass2.png
│       ├── bass3.png
│       ├── bass4.png
│       ├── bass_clef.png
│       ├── bass_full.png
│       ├── beats.png
│       ├── bin.png
│       ├── bongos.png
│       ├── cabasa.png
│       ├── cello.png
│       ├── congas.png
│       ├── cowbell.png
│       ├── cowbell_more.png
│       ├── cymbal_large.png
│       ├── cymbal_small.png
│       ├── deck.png
│       ├── double_bass.png
│       ├── drumbox.png
│       ├── drums.png
│       ├── envelope.png
│       ├── female.png
│       ├── female_head.png
│       ├── ff.png
│       ├── film.png
│       ├── folder.png
│       ├── folder_down.png
│       ├── folder_left.png
│       ├── folder_right.png
│       ├── folder_up.png
│       ├── fx.png
│       ├── group.png
│       ├── guitar.png
│       ├── guitar2.png
│       ├── guitar3.png
│       ├── guitar4.png
│       ├── guitar5.png
│       ├── guitar_full.png
│       ├── harmonica.png
│       ├── harp.png
│       ├── hihat.png
│       ├── idea.png
│       ├── kick.png
│       ├── male.png
│       ├── male_head.png
│       ├── maracas.png
│       ├── meter.png
│       ├── mic.png
│       ├── mic_condenser_1.png
│       ├── mic_condenser_2.png
│       ├── mic_dynamic_1.png
│       ├── mic_dynamic_2.png
│       ├── mic_shotgun.png
│       ├── midi.png
│       ├── mixer.png
│       ├── organ.png
│       ├── overheads.png
│       ├── pads.png
│       ├── pedal.png
│       ├── phones.png
│       ├── piano.png
│       ├── pp.png
│       ├── reverb.png
│       ├── ride_bell.png
│       ├── ride_rim.png
│       ├── room_large.png
│       ├── room_medium.png
│       ├── room_small.png
│       ├── sax.png
│       ├── snare_bottom.png
│       ├── snare_top.png
│       ├── speech.png
│       ├── synth.png
│       ├── synth2.png
│       ├── synthbass.png
│       ├── system.png
│       ├── tamborine.png
│       ├── tape.png
│       ├── tom.png
│       ├── treble_clef.png
│       ├── trombone.png
│       ├── trumpet.png
│       ├── violin.png
│       ├── xylophone.png
│       └── yeah_you_guys_are_great.png
├── Effects
│   ├── HeDa
│   │   └── Track_Inspector
│   ├── IX
│   │   ├── MIDI_CCRider
│   │   ├── MIDI_DuplicateFilter
│   │   ├── MIDI_KeyMap
│   │   ├── MIDI_KeySnap
│   │   ├── MIDI_Router
│   │   ├── MIDI_Tool
│   │   ├── MIDI_Tool II
│   │   ├── MIDI_Variant
│   │   ├── MIDI_Velocifier II
│   │   ├── MIDI_Wobulator
│   │   ├── Mixer_8xM-1xS
│   │   ├── Mixer_8xS-1xS
│   │   ├── PhaseAdjustingRouter
│   │   ├── StereoPhaseInvert
│   │   ├── Switcher2
│   │   └── SwixMitch
│   ├── Liteon
│   │   ├── 3bandpeakfilter
│   │   ├── applefilter12db
│   │   ├── applefilter72db
│   │   ├── bassmanager
│   │   ├── butterworth24db
│   │   ├── cheby24db
│   │   ├── deesser
│   │   ├── lorenzattractor
│   │   ├── moog24db
│   │   ├── nonlinear
│   │   ├── np1136peaklimiter
│   │   ├── pinknoisegen
│   │   ├── presenceeq
│   │   ├── pseudostereo
│   │   ├── rbjstereofilter12db
│   │   ├── ringmodulator
│   │   ├── shelvingfilter
│   │   ├── simplelp6db
│   │   ├── statevariable
│   │   ├── tilteq
│   │   ├── vumetergfx
│   │   ├── vumetergfxsum
│   │   └── waveshapermulti
│   ├── Teej
│   │   ├── rbj12eq-teej
│   │   ├── rbj4eq-teej
│   │   └── rbj4notch-teej
│   ├── Till
│   │   ├── Auto-Wideness v1.0
│   │   ├── Transient-driven Auto-Pan v1.0 (Receiver)
│   │   ├── Transient-driven Auto-Pan v1.0 (Transmitter)
│   │   ├── Transient-driven Auto-Pan v1.1 (Receiver)
│   │   └── Transient-driven Auto-Pan v1.1 (Transmitter)
│   ├── analysis
│   │   ├── WigMCVUMeter
│   │   ├── gain_reduction_scope
│   │   ├── gfxanalyzer
│   │   ├── gfxscope
│   │   ├── gfxspectrograph
│   │   ├── loudness_meter
│   │   ├── spectropaint
│   │   └── zoomanalyzer
│   ├── delay
│   │   ├── delay
│   │   ├── delay_chfun
│   │   ├── delay_chorus
│   │   ├── delay_lowres
│   │   ├── delay_sustain
│   │   ├── delay_tone
│   │   ├── delay_varlen
│   │   └── time_adjustment
│   ├── dynamics
│   │   ├── general_dynamics
│   │   └── general_dynamics_track_inspector
│   ├── filters
│   │   ├── autopeakfilter
│   │   ├── dc_remove
│   │   ├── mdct-filter
│   │   ├── mdct-volsweep
│   │   ├── resonantlowpass
│   │   ├── spectro_filter_paint
│   │   └── sweepinglowpass
│   ├── guitar
│   │   ├── amp-model
│   │   ├── amp-model-dual
│   │   ├── chorus
│   │   ├── distort-fuzz
│   │   ├── distortion
│   │   ├── flanger
│   │   ├── phaser
│   │   ├── tremolo
│   │   └── wah
│   ├── loopsamplers
│   │   ├── autoloop
│   │   ├── loopsampler-granul
│   │   ├── loopsampler-m2
│   │   ├── loopsampler-m2-midi
│   │   └── super8
│   ├── loser
│   │   ├── 3BandEQ
│   │   ├── 3BandJoiner
│   │   ├── 3BandSplitter
│   │   ├── 4BandEQ
│   │   ├── 4BandJoiner
│   │   ├── 4BandSplitter
│   │   ├── 50HzKicker
│   │   ├── 5BandJoiner
│   │   ├── 5BandSplitter
│   │   ├── CenterCanceler
│   │   ├── Compciter
│   │   ├── DDC
│   │   ├── DVC2JS
│   │   ├── DVCJS
│   │   ├── Exciter
│   │   ├── MGA_JSLimiter
│   │   ├── MGA_JSLimiterST
│   │   ├── MIDI_EQ_Ducker
│   │   ├── SP1LimiterJS
│   │   ├── Saturation
│   │   ├── StereoField
│   │   ├── TimeDelayer
│   │   ├── TransientController
│   │   ├── TransientKiller
│   │   ├── UpwardExpander
│   │   ├── WhiteNoise
│   │   ├── ZeroCrossingMaximizer
│   │   ├── amplitudeModulator
│   │   ├── gfxGoniometer
│   │   ├── masterLimiter
│   │   ├── phaseMeter
│   │   ├── ppp
│   │   ├── stereo
│   │   ├── stereoEnhancer
│   │   ├── timeDifferencePan
│   │   └── waveShapingDstr
│   ├── meters
│   │   ├── bit_meter
│   │   └── dynamics_meter
│   ├── midi
│   │   ├── drumtrigger
│   │   ├── midi_CC_mapper
│   │   ├── midi_arp
│   │   ├── midi_choke
│   │   ├── midi_chorderizer
│   │   ├── midi_chordkey
│   │   ├── midi_delay
│   │   ├── midi_logger
│   │   ├── midi_maptokey
│   │   ├── midi_note2channel
│   │   ├── midi_note_filter
│   │   ├── midi_note_hold
│   │   ├── midi_note_repeater
│   │   ├── midi_note_sanitizer
│   │   ├── midi_transpose
│   │   ├── midi_velocitycontrol
│   │   ├── midinoteondelay
│   │   ├── mtc_logger
│   │   ├── program_bank_onload
│   │   ├── sequencer_baby
│   │   ├── sequencer_baby_v2
│   │   └── sequencer_megababy
│   ├── misc
│   │   ├── adpcm_simulator
│   │   ├── reverseness
│   │   ├── tonifier
│   │   ├── tonifier2
│   │   └── video_sample_peeker
│   ├── nvk-ReaScripts
│   │   ├── AUTODOPPLER
│   │   │   └── nvk_DOPPLER.jsfx
│   │   └── CREATE
│   │       └── nvk_RECORD.jsfx
│   ├── old_unsupported
│   │   └── dither
│   ├── pitch
│   │   ├── mdct-shift
│   │   ├── octavedown
│   │   ├── octaveup
│   │   ├── pitch2
│   │   ├── pitchdown
│   │   └── superpitch
│   ├── remaincalm_org
│   │   ├── avocado_glitch
│   │   ├── floaty_delay
│   │   ├── paranoia_mangler
│   │   ├── paranoia_mangler.rpl
│   │   └── tonegate
│   ├── schwa
│   │   ├── audio_statistics
│   │   ├── fft_splitter
│   │   ├── gaussian_noise_generator
│   │   ├── midi_examine
│   │   ├── midi_humanizer
│   │   ├── midi_modal_randomness
│   │   └── soft_clipper
│   ├── sstillwell
│   │   ├── 1175
│   │   ├── 3x3
│   │   ├── 3x3_6dbSlope
│   │   ├── 4x4
│   │   ├── autoexpand
│   │   ├── badbussmojo
│   │   ├── badbussmojo_aa
│   │   ├── chorus
│   │   ├── chorus_stereo
│   │   ├── delay_pong
│   │   ├── delay_tempo
│   │   ├── dirtsqueeze
│   │   ├── eventhorizon
│   │   ├── eventhorizon2
│   │   ├── exciter
│   │   ├── expander
│   │   ├── expressbus
│   │   ├── fairlychildish
│   │   ├── fairlychildish2
│   │   ├── flangebaby
│   │   ├── hpflpf
│   │   ├── hugebooty
│   │   ├── louderizer
│   │   ├── louderizer_lpf
│   │   ├── majortom
│   │   ├── mastertom
│   │   ├── ozzifier
│   │   ├── randomizer
│   │   ├── rbj1073
│   │   ├── rbj4eq
│   │   ├── rbj7eq
│   │   ├── realoud
│   │   ├── realoud_lpf
│   │   ├── stereowidth
│   │   ├── thunderkick
│   │   ├── volscale
│   │   └── width
│   ├── synthesis
│   │   ├── sine_sweep
│   │   ├── spectral_hold
│   │   └── tonegenerator
│   ├── utility
│   │   ├── KanakaMS5
│   │   ├── KanakaMSEncoder1
│   │   ├── chanmix2
│   │   ├── channel_mapper
│   │   ├── channelmixer
│   │   ├── dither_psycho
│   │   ├── phase_adjust
│   │   ├── smpte_ltc_reader
│   │   ├── vca_follow
│   │   ├── vca_lead
│   │   ├── volume
│   │   ├── volume_pan
│   │   └── volume_pan_sample_accurate_auto
│   └── waveshapers
│       ├── graphdist
│       └── graphdist-dyn
├── FXChains
│   └── Track Inspector Monitor FX.RfxChain
├── KeyMaps
│   ├── DK keymap.ReaperKeyMap
│   └── German Keymap.ReaperKeyMap
├── LangPack
├── MIDINoteNames
│   └── note_name_sample.txt
├── MetadataCaches
│   └── dc
│       └── 01a9b19d8af276fea3ed5f766f917d4e30f96c.MetadataCache
├── MouseMaps
├── OSC
│   ├── Default.ReaperOSC
│   ├── LogicPad.ReaperOSC
│   └── LogicTouch.ReaperOSC
├── ProjectTemplates
├── ReaImGui
│   └── 02CB7EDB.ini
├── ReaPack
│   ├── cache
│   │   ├── MPL Scripts.xml
│   │   ├── ReaPack.xml
│   │   ├── ReaTeam Extensions.xml
│   │   ├── ReaTeam JSFX.xml
│   │   ├── ReaTeam LangPacks.xml
│   │   ├── ReaTeam Scripts.xml
│   │   ├── ReaTeam Themes.xml
│   │   ├── X-Raym Scripts.xml
│   │   └── nvk-ReaScripts.xml
│   └── registry.db
├── S&M.ini
├── Scripts
│   ├── Cockos
│   │   ├── Default_6.0_theme_adjuster.lua
│   │   ├── Default_6.0_theme_adjuster_images
│   │   │   ├── apply_100.png
│   │   │   ├── apply_100@2x.png
│   │   │   ├── apply_150.png
│   │   │   ├── apply_150@2x.png
│   │   │   ├── apply_200.png
│   │   │   ├── apply_200@2x.png
│   │   │   ├── apply_50.png
│   │   │   ├── apply_50@2x.png
│   │   │   ├── apply_75.png
│   │   │   ├── apply_75@2x.png
│   │   │   ├── apply_to_sel.png
│   │   │   ├── apply_to_sel@2x.png
│   │   │   ├── apply_to_sel_docked.png
│   │   │   ├── apply_to_sel_docked@2x.png
│   │   │   ├── bin.png
│   │   │   ├── bin@2x.png
│   │   │   ├── button_empty.png
│   │   │   ├── button_empty@2x.png
│   │   │   ├── button_left.png
│   │   │   ├── button_left@2x.png
│   │   │   ├── button_left_dark.png
│   │   │   ├── button_left_dark@2x.png
│   │   │   ├── button_right.png
│   │   │   ├── button_right@2x.png
│   │   │   ├── button_right_dark.png
│   │   │   ├── button_right_dark@2x.png
│   │   │   ├── cell_hide.png
│   │   │   ├── cell_hide@2x.png
│   │   │   ├── cell_hide_all.png
│   │   │   ├── cell_hide_all@2x.png
│   │   │   ├── cell_hide_on.png
│   │   │   ├── cell_hide_on@2x.png
│   │   │   ├── cell_tick.png
│   │   │   ├── cell_tick@2x.png
│   │   │   ├── cell_tick_on.png
│   │   │   ├── cell_tick_on@2x.png
│   │   │   ├── center_transport.png
│   │   │   ├── center_transport@2x.png
│   │   │   ├── center_transport_on.png
│   │   │   ├── center_transport_on@2x.png
│   │   │   ├── color_apply.png
│   │   │   ├── color_apply@2x.png
│   │   │   ├── color_apply_all.png
│   │   │   ├── color_apply_all@2x.png
│   │   │   ├── color_apply_all_on.png
│   │   │   ├── color_apply_all_on@2x.png
│   │   │   ├── color_apply_large.png
│   │   │   ├── color_apply_large@2x.png
│   │   │   ├── color_dim.png
│   │   │   ├── color_dim@2x.png
│   │   │   ├── color_dim_all.png
│   │   │   ├── color_dim_all@2x.png
│   │   │   ├── color_text.png
│   │   │   ├── color_text@2x.png
│   │   │   ├── color_text_on.png
│   │   │   ├── color_text_on@2x.png
│   │   │   ├── dock.png
│   │   │   ├── dock@2x.png
│   │   │   ├── docked_edit.png
│   │   │   ├── docked_edit@2x.png
│   │   │   ├── faderbg_gamma.png
│   │   │   ├── faderbg_gamma@2x.png
│   │   │   ├── faderbg_saturation.png
│   │   │   ├── faderbg_saturation@2x.png
│   │   │   ├── faderbg_tint.png
│   │   │   ├── faderbg_tint@2x.png
│   │   │   ├── help.png
│   │   │   ├── help@2x.png
│   │   │   ├── helpHeader_l.png
│   │   │   ├── helpHeader_l@2x.png
│   │   │   ├── helpHeader_r.png
│   │   │   ├── helpHeader_r@2x.png
│   │   │   ├── help_on.png
│   │   │   ├── help_on@2x.png
│   │   │   ├── icon_warning_on.png
│   │   │   ├── icon_warning_on@2x.png
│   │   │   ├── key_env.png
│   │   │   ├── key_env@2x.png
│   │   │   ├── key_fx.png
│   │   │   ├── key_fx@2x.png
│   │   │   ├── key_mon.png
│   │   │   ├── key_mon@2x.png
│   │   │   ├── key_pan.png
│   │   │   ├── key_pan@2x.png
│   │   │   ├── key_recarm.png
│   │   │   ├── key_recarm@2x.png
│   │   │   ├── key_recinput.png
│   │   │   ├── key_recinput@2x.png
│   │   │   ├── key_recmode.png
│   │   │   ├── key_recmode@2x.png
│   │   │   ├── key_route.png
│   │   │   ├── key_route@2x.png
│   │   │   ├── key_vol.png
│   │   │   ├── key_vol@2x.png
│   │   │   ├── layout_A.png
│   │   │   ├── layout_A@2x.png
│   │   │   ├── layout_A_on.png
│   │   │   ├── layout_A_on@2x.png
│   │   │   ├── layout_B.png
│   │   │   ├── layout_B@2x.png
│   │   │   ├── layout_B_on.png
│   │   │   ├── layout_B_on@2x.png
│   │   │   ├── layout_C.png
│   │   │   ├── layout_C@2x.png
│   │   │   ├── layout_C_on.png
│   │   │   ├── layout_C_on@2x.png
│   │   │   ├── layout_docked_A.png
│   │   │   ├── layout_docked_A@2x.png
│   │   │   ├── layout_docked_A_on.png
│   │   │   ├── layout_docked_A_on@2x.png
│   │   │   ├── layout_docked_B.png
│   │   │   ├── layout_docked_B@2x.png
│   │   │   ├── layout_docked_B_on.png
│   │   │   ├── layout_docked_B_on@2x.png
│   │   │   ├── layout_docked_C.png
│   │   │   ├── layout_docked_C@2x.png
│   │   │   ├── layout_docked_C_on.png
│   │   │   ├── layout_docked_C_on@2x.png
│   │   │   ├── left.png
│   │   │   ├── left@2x.png
│   │   │   ├── match_folder_indent.png
│   │   │   ├── match_folder_indent@2x.png
│   │   │   ├── match_folder_indent_on.png
│   │   │   ├── match_folder_indent_on@2x.png
│   │   │   ├── multi_row.png
│   │   │   ├── multi_row@2x.png
│   │   │   ├── multi_row_on.png
│   │   │   ├── multi_row_on@2x.png
│   │   │   ├── next_prev.png
│   │   │   ├── next_prev@2x.png
│   │   │   ├── next_prev_on.png
│   │   │   ├── next_prev_on@2x.png
│   │   │   ├── page_titles.png
│   │   │   ├── page_titles@2x.png
│   │   │   ├── page_titles_small.png
│   │   │   ├── page_titles_small@2x.png
│   │   │   ├── play_text.png
│   │   │   ├── play_text@2x.png
│   │   │   ├── play_text_on.png
│   │   │   ├── play_text_on@2x.png
│   │   │   ├── right.png
│   │   │   ├── right@2x.png
│   │   │   ├── scroll_sel.png
│   │   │   ├── scroll_sel@2x.png
│   │   │   ├── scroll_sel_on.png
│   │   │   ├── scroll_sel_on@2x.png
│   │   │   ├── show_fx.png
│   │   │   ├── show_fx@2x.png
│   │   │   ├── show_fx_on.png
│   │   │   ├── show_fx_on@2x.png
│   │   │   ├── show_icons.png
│   │   │   ├── show_icons@2x.png
│   │   │   ├── show_icons_on.png
│   │   │   ├── show_icons_on@2x.png
│   │   │   ├── show_param.png
│   │   │   ├── show_param@2x.png
│   │   │   ├── show_param_on.png
│   │   │   ├── show_param_on@2x.png
│   │   │   ├── show_play_rate.png
│   │   │   ├── show_play_rate@2x.png
│   │   │   ├── show_play_rate_on.png
│   │   │   ├── show_play_rate_on@2x.png
│   │   │   ├── show_send.png
│   │   │   ├── show_send@2x.png
│   │   │   ├── show_send_on.png
│   │   │   ├── show_send_on@2x.png
│   │   │   ├── slider.png
│   │   │   ├── slider@2x.png
│   │   │   ├── time_sig.png
│   │   │   ├── time_sig@2x.png
│   │   │   ├── time_sig_on.png
│   │   │   ├── time_sig_on@2x.png
│   │   │   ├── transport_title.png
│   │   │   └── transport_title@2x.png
│   │   ├── Default_7.0_theme_adjuster.lua
│   │   ├── Default_7.0_theme_adjuster_images
│   │   │   ├── always_hide.png
│   │   │   ├── always_hide_150.png
│   │   │   ├── always_hide_200.png
│   │   │   ├── always_hide_on.png
│   │   │   ├── always_hide_on_150.png
│   │   │   ├── always_hide_on_200.png
│   │   │   ├── apply_colour_all.png
│   │   │   ├── apply_colour_all_150.png
│   │   │   ├── apply_colour_all_200.png
│   │   │   ├── apply_colour_all_on.png
│   │   │   ├── apply_colour_all_on_150.png
│   │   │   ├── apply_colour_all_on_200.png
│   │   │   ├── arcStalk.png
│   │   │   ├── arcStalk_150.png
│   │   │   ├── arcStalk_200.png
│   │   │   ├── arcStalk_closed.png
│   │   │   ├── arcStalk_closed_150.png
│   │   │   ├── arcStalk_closed_200.png
│   │   │   ├── bentNail.png
│   │   │   ├── bentNail_150.png
│   │   │   ├── bentNail_200.png
│   │   │   ├── bin.png
│   │   │   ├── bin_150.png
│   │   │   ├── bin_200.png
│   │   │   ├── button_leftArrow_off.png
│   │   │   ├── button_leftArrow_off_150.png
│   │   │   ├── button_leftArrow_off_200.png
│   │   │   ├── button_leftArrow_on.png
│   │   │   ├── button_leftArrow_on_150.png
│   │   │   ├── button_leftArrow_on_200.png
│   │   │   ├── button_off.png
│   │   │   ├── button_off_150.png
│   │   │   ├── button_off_200.png
│   │   │   ├── button_on.png
│   │   │   ├── button_on_150.png
│   │   │   ├── button_on_200.png
│   │   │   ├── button_upArrow_off.png
│   │   │   ├── button_upArrow_off_150.png
│   │   │   ├── button_upArrow_off_200.png
│   │   │   ├── button_upArrow_on.png
│   │   │   ├── button_upArrow_on_150.png
│   │   │   ├── button_upArrow_on_200.png
│   │   │   ├── col_chooser.png
│   │   │   ├── col_chooser_150.png
│   │   │   ├── col_chooser_200.png
│   │   │   ├── color_apply.png
│   │   │   ├── color_apply_150.png
│   │   │   ├── color_apply_200.png
│   │   │   ├── drag_marker.png
│   │   │   ├── export.png
│   │   │   ├── export_150.png
│   │   │   ├── export_200.png
│   │   │   ├── export_all.png
│   │   │   ├── export_all_150.png
│   │   │   ├── export_all_200.png
│   │   │   ├── folderBalance_all.png
│   │   │   ├── folderBalance_all_150.png
│   │   │   ├── folderBalance_all_200.png
│   │   │   ├── folderBalance_none.png
│   │   │   ├── folderBalance_none_150.png
│   │   │   ├── folderBalance_none_200.png
│   │   │   ├── folderBalance_stretch.png
│   │   │   ├── folderBalance_stretch_150.png
│   │   │   ├── folderBalance_stretch_200.png
│   │   │   ├── font_size_stack.png
│   │   │   ├── font_size_stack_150.png
│   │   │   ├── font_size_stack_200.png
│   │   │   ├── halfCapsule_nonInteractive.png
│   │   │   ├── halfCapsule_nonInteractive_150.png
│   │   │   ├── halfCapsule_nonInteractive_200.png
│   │   │   ├── hide.png
│   │   │   ├── hide_150.png
│   │   │   ├── hide_200.png
│   │   │   ├── import.png
│   │   │   ├── import_150.png
│   │   │   ├── import_200.png
│   │   │   ├── input_v.png
│   │   │   ├── input_v_150.png
│   │   │   ├── input_v_200.png
│   │   │   ├── knobStack_20px_dark.png
│   │   │   ├── knobStack_20px_dark_150.png
│   │   │   ├── knobStack_20px_dark_200.png
│   │   │   ├── layoutA_off.png
│   │   │   ├── layoutA_off_150.png
│   │   │   ├── layoutA_off_200.png
│   │   │   ├── layoutA_on.png
│   │   │   ├── layoutA_on_150.png
│   │   │   ├── layoutA_on_200.png
│   │   │   ├── layoutB_off.png
│   │   │   ├── layoutB_off_150.png
│   │   │   ├── layoutB_off_200.png
│   │   │   ├── layoutB_on.png
│   │   │   ├── layoutB_on_150.png
│   │   │   ├── layoutB_on_200.png
│   │   │   ├── layoutC_off.png
│   │   │   ├── layoutC_off_150.png
│   │   │   ├── layoutC_off_200.png
│   │   │   ├── layoutC_on.png
│   │   │   ├── layoutC_on_150.png
│   │   │   ├── layoutC_on_200.png
│   │   │   ├── mcp_env.png
│   │   │   ├── mcp_env_150.png
│   │   │   ├── mcp_env_200.png
│   │   │   ├── mcp_fx.png
│   │   │   ├── mcp_fx_150.png
│   │   │   ├── mcp_fx_200.png
│   │   │   ├── mcp_io.png
│   │   │   ├── mcp_io_150.png
│   │   │   ├── mcp_io_200.png
│   │   │   ├── mcp_mutesolo.png
│   │   │   ├── mcp_mutesolo_150.png
│   │   │   ├── mcp_mutesolo_200.png
│   │   │   ├── mcp_phase.png
│   │   │   ├── mcp_phase_150.png
│   │   │   ├── mcp_phase_200.png
│   │   │   ├── mcp_recmon.png
│   │   │   ├── mcp_recmon_150.png
│   │   │   ├── mcp_recmon_200.png
│   │   │   ├── mcp_volthumb.png
│   │   │   ├── mcp_volthumb_150.png
│   │   │   ├── mcp_volthumb_200.png
│   │   │   ├── mcpstrip_input.png
│   │   │   ├── mcpstrip_input_150.png
│   │   │   ├── mcpstrip_input_200.png
│   │   │   ├── mcpstrip_panwidth.png
│   │   │   ├── mcpstrip_panwidth_150.png
│   │   │   ├── mcpstrip_panwidth_200.png
│   │   │   ├── monitorScale.png
│   │   │   ├── monitorScale_150.png
│   │   │   ├── monitorScale_200.png
│   │   │   ├── page_envcp.png
│   │   │   ├── page_envcp_150.png
│   │   │   ├── page_envcp_200.png
│   │   │   ├── page_error.png
│   │   │   ├── page_error_150.png
│   │   │   ├── page_error_200.png
│   │   │   ├── page_global.png
│   │   │   ├── page_global_150.png
│   │   │   ├── page_global_200.png
│   │   │   ├── page_mcp.png
│   │   │   ├── page_mcp_150.png
│   │   │   ├── page_mcp_200.png
│   │   │   ├── page_ol_off.png
│   │   │   ├── page_ol_off_150.png
│   │   │   ├── page_ol_off_200.png
│   │   │   ├── page_ol_on.png
│   │   │   ├── page_ol_on_150.png
│   │   │   ├── page_ol_on_200.png
│   │   │   ├── page_tcp.png
│   │   │   ├── page_tcp_150.png
│   │   │   ├── page_tcp_200.png
│   │   │   ├── page_theme.png
│   │   │   ├── page_theme_150.png
│   │   │   ├── page_theme_200.png
│   │   │   ├── page_trans.png
│   │   │   ├── page_trans_150.png
│   │   │   ├── page_trans_200.png
│   │   │   ├── rect20_nonInteractive.png
│   │   │   ├── rect20_nonInteractive_150.png
│   │   │   ├── rect20_nonInteractive_200.png
│   │   │   ├── refresh.png
│   │   │   ├── refresh_150.png
│   │   │   ├── refresh_200.png
│   │   │   ├── round20_nonInteractive.png
│   │   │   ├── round20_nonInteractive_150.png
│   │   │   ├── round20_nonInteractive_200.png
│   │   │   ├── round22_nonInteractive.png
│   │   │   ├── round22_nonInteractive_150.png
│   │   │   ├── round22_nonInteractive_200.png
│   │   │   ├── roundBox_white25.png
│   │   │   ├── roundBox_white25_150.png
│   │   │   ├── roundBox_white25_200.png
│   │   │   ├── scrollbar_v.png
│   │   │   ├── scrollbar_v_150.png
│   │   │   ├── scrollbar_v_200.png
│   │   │   ├── selDotOff.png
│   │   │   ├── selDotOff_150.png
│   │   │   ├── selDotOff_200.png
│   │   │   ├── selDotSel.png
│   │   │   ├── selDotSel_150.png
│   │   │   ├── selDotSel_200.png
│   │   │   ├── selDotUnsel.png
│   │   │   ├── selDotUnsel_150.png
│   │   │   ├── selDotUnsel_200.png
│   │   │   ├── show.png
│   │   │   ├── show_150.png
│   │   │   ├── show_200.png
│   │   │   ├── slider.png
│   │   │   ├── slider_150.png
│   │   │   ├── slider_200.png
│   │   │   ├── slider_nonInteractive.png
│   │   │   ├── slider_nonInteractive_150.png
│   │   │   ├── slider_nonInteractive_200.png
│   │   │   ├── spinner_left.png
│   │   │   ├── spinner_left_150.png
│   │   │   ├── spinner_left_200.png
│   │   │   ├── swatch.png
│   │   │   ├── swatch_150.png
│   │   │   ├── swatch_200.png
│   │   │   ├── tcpDrop_nonInteractive.png
│   │   │   ├── tcpDrop_nonInteractive_150.png
│   │   │   ├── tcpDrop_nonInteractive_200.png
│   │   │   ├── tcpInFx_nonInteractive.png
│   │   │   ├── tcpInFx_nonInteractive_150.png
│   │   │   ├── tcpInFx_nonInteractive_200.png
│   │   │   ├── tcp_Section_disable_off.png
│   │   │   ├── tcp_Section_disable_off_150.png
│   │   │   ├── tcp_Section_disable_off_200.png
│   │   │   ├── tcp_Section_disable_on.png
│   │   │   ├── tcp_Section_disable_on_150.png
│   │   │   ├── tcp_Section_disable_on_200.png
│   │   │   ├── tcp_Section_pin_off.png
│   │   │   ├── tcp_Section_pin_off_150.png
│   │   │   ├── tcp_Section_pin_off_200.png
│   │   │   ├── tcp_Section_pin_on.png
│   │   │   ├── tcp_Section_pin_on_150.png
│   │   │   ├── tcp_Section_pin_on_200.png
│   │   │   ├── tcp_embedUi.png
│   │   │   ├── tcp_embedUi_150.png
│   │   │   ├── tcp_embedUi_200.png
│   │   │   ├── tcp_env.png
│   │   │   ├── tcp_env_150.png
│   │   │   ├── tcp_env_200.png
│   │   │   ├── tcp_fx.png
│   │   │   ├── tcp_fx_150.png
│   │   │   ├── tcp_fx_200.png
│   │   │   ├── tcp_infx.png
│   │   │   ├── tcp_infx_150.png
│   │   │   ├── tcp_infx_200.png
│   │   │   ├── tcp_insert.png
│   │   │   ├── tcp_insert_150.png
│   │   │   ├── tcp_insert_200.png
│   │   │   ├── tcp_io.png
│   │   │   ├── tcp_io_150.png
│   │   │   ├── tcp_io_200.png
│   │   │   ├── tcp_monitor.png
│   │   │   ├── tcp_monitor_150.png
│   │   │   ├── tcp_monitor_200.png
│   │   │   ├── tcp_mutesolo.png
│   │   │   ├── tcp_mutesolo_150.png
│   │   │   ├── tcp_mutesolo_200.png
│   │   │   ├── tcp_panwidth.png
│   │   │   ├── tcp_panwidth_150.png
│   │   │   ├── tcp_panwidth_200.png
│   │   │   ├── tcp_param.png
│   │   │   ├── tcp_param_150.png
│   │   │   ├── tcp_param_200.png
│   │   │   ├── tcp_phase.png
│   │   │   ├── tcp_phase_150.png
│   │   │   ├── tcp_phase_200.png
│   │   │   ├── tcp_recarm.png
│   │   │   ├── tcp_recarm_150.png
│   │   │   ├── tcp_recarm_200.png
│   │   │   ├── tcp_recinput.png
│   │   │   ├── tcp_recinput_150.png
│   │   │   ├── tcp_recinput_200.png
│   │   │   ├── tcp_recmode.png
│   │   │   ├── tcp_recmode_150.png
│   │   │   ├── tcp_recmode_200.png
│   │   │   ├── tcp_send.png
│   │   │   ├── tcp_send_150.png
│   │   │   ├── tcp_send_200.png
│   │   │   ├── tcp_vol.png
│   │   │   ├── tcp_vol_150.png
│   │   │   ├── tcp_vol_200.png
│   │   │   ├── tcp_vol_knob_stack.png
│   │   │   ├── tcp_vol_knob_stack_150.png
│   │   │   ├── tcp_vol_knob_stack_200.png
│   │   │   ├── transSectionBg.png
│   │   │   ├── transSectionBg_150.png
│   │   │   ├── transSectionBg_200.png
│   │   │   ├── vEdge_resize.png
│   │   │   ├── vEdge_resize_150.png
│   │   │   └── vEdge_resize_200.png
│   │   └── lyrics.lua
│   ├── HeDaScripts
│   │   ├── Auto Record Mode
│   │   │   ├── Auto Record Mode - changelog.txt
│   │   │   ├── Auto Record Mode.dat
│   │   │   ├── HeDa_Auto Record Mode.eel
│   │   │   └── appicon.png
│   │   ├── HeDaMixer
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Default themes
│   │   │   │   └── Default_5.0
│   │   │   │       ├── BEBAS.ttf
│   │   │   │       ├── Default_5.0.hm_theme
│   │   │   │       ├── faders
│   │   │   │       │   ├── Default.png
│   │   │   │       │   ├── Fader Blue.png
│   │   │   │       │   ├── Fader Cyan.png
│   │   │   │       │   ├── Fader Gold.png
│   │   │   │       │   ├── Fader Green.png
│   │   │   │       │   ├── Fader Pink.png
│   │   │   │       │   ├── Fader Red.png
│   │   │   │       │   ├── Fader Violet.png
│   │   │   │       │   ├── Fader3.0.png
│   │   │   │       │   ├── Fader4.0.png
│   │   │   │       │   └── Small Fader.png
│   │   │   │       ├── mcp_io.png
│   │   │   │       ├── mcp_mute_off.png
│   │   │   │       ├── mcp_mute_on.png
│   │   │   │       ├── mcp_solo_off.png
│   │   │   │       ├── mcp_solo_on.png
│   │   │   │       ├── panfaders
│   │   │   │       │   ├── Default.png
│   │   │   │       │   ├── Fader Blue.png
│   │   │   │       │   ├── Fader Cyan.png
│   │   │   │       │   ├── Fader Gold.png
│   │   │   │       │   ├── Fader Green.png
│   │   │   │       │   ├── Fader Pink.png
│   │   │   │       │   ├── Fader Red.png
│   │   │   │       │   ├── Fader Violet.png
│   │   │   │       │   ├── Fader3.0.png
│   │   │   │       │   ├── Fader4.0.png
│   │   │   │       │   └── Small Fader.png
│   │   │   │       ├── track_fx_byp.png
│   │   │   │       ├── track_fx_empty.png
│   │   │   │       └── track_fx_norm.png
│   │   │   ├── HMx32.dat
│   │   │   ├── HMx64.dat
│   │   │   ├── HeDaMixer - changelog.txt
│   │   │   ├── HeDaMixer.dat
│   │   │   ├── HeDaMixer.lua
│   │   │   ├── License.txt
│   │   │   └── appicon.png
│   │   ├── HeDaMixer settings
│   │   ├── HeDaScripts Manager
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Global Options.ini
│   │   │   ├── HSMux32.dat
│   │   │   ├── HSMux32_7.dat
│   │   │   ├── HSMux64.dat
│   │   │   ├── HSMux64_7.dat
│   │   │   ├── HSMx32.dat
│   │   │   ├── HSMx32_7.dat
│   │   │   ├── HSMx64.dat
│   │   │   ├── HSMx64_7.dat
│   │   │   ├── HeDaScripts Auto-check updates.lua
│   │   │   ├── HeDaScripts Manager - changelog.txt
│   │   │   ├── HeDaScripts Manager.dat
│   │   │   ├── HeDaScripts Manager.lua
│   │   │   ├── Installation Instructions.txt
│   │   │   ├── License.txt
│   │   │   ├── _resources
│   │   │   │   ├── 7z
│   │   │   │   │   ├── 7zz
│   │   │   │   │   └── License.txt
│   │   │   │   ├── Archiving - appicon.png
│   │   │   │   ├── Auto Bypass - appicon.png
│   │   │   │   ├── Color - appicon.png
│   │   │   │   ├── HeDaMixer Graph VIP - appicon.png
│   │   │   │   ├── HeDaMixer VIP - appicon.png
│   │   │   │   ├── Import Export Subtitles - appicon.png
│   │   │   │   ├── Item Timecode - appicon.png
│   │   │   │   ├── Loudness Graph VIP - appicon.png
│   │   │   │   ├── MIDI Dynamics VIP - appicon.png
│   │   │   │   ├── Notes Reader VIP - appicon.png
│   │   │   │   ├── Panel VIP - appicon.png
│   │   │   │   ├── Projects VIP - appicon.png
│   │   │   │   ├── ReaWhisper - appicon.png
│   │   │   │   ├── Rec Arm - appicon.png
│   │   │   │   ├── Region Tracks VIP - appicon.png
│   │   │   │   ├── Render LUFS VIP - appicon.png
│   │   │   │   ├── Render Loop - appicon.png
│   │   │   │   ├── Solo Tracks of Selected Items - appicon.png
│   │   │   │   ├── Subliminal - appicon.png
│   │   │   │   ├── Track Attributes - appicon.png
│   │   │   │   ├── Track Inspector 2 VIP - appicon.png
│   │   │   │   ├── Track Inspector VIP - appicon.png
│   │   │   │   ├── Track Templates - appicon.png
│   │   │   │   ├── Track Templates VIP - appicon.png
│   │   │   │   ├── Tracks - appicon.png
│   │   │   │   ├── Transpose - appicon.png
│   │   │   │   ├── Versions - appicon.png
│   │   │   │   └── logo.png
│   │   │   └── appicon.png
│   │   ├── HeDaScripts Manager settings
│   │   │   ├── 7z
│   │   │   │   ├── 7zz
│   │   │   │   └── License.txt
│   │   │   ├── _Settings.cfg
│   │   │   └── _resources
│   │   ├── HeDaScripts Manager.log
│   │   ├── Import Export TXT
│   │   │   ├── HeDa_Export Text Items to TXT with timestamps.lua
│   │   │   ├── HeDa_Export Text Items to TXT x32.dat
│   │   │   ├── HeDa_Export Text Items to TXT x32_7.dat
│   │   │   ├── HeDa_Export Text Items to TXT x64.dat
│   │   │   ├── HeDa_Export Text Items to TXT x64_7.dat
│   │   │   ├── HeDa_Export Text Items to TXT.lua
│   │   │   ├── HeDa_Import TXT to Text Items x32.dat
│   │   │   ├── HeDa_Import TXT to Text Items x32_7.dat
│   │   │   ├── HeDa_Import TXT to Text Items x64.dat
│   │   │   ├── HeDa_Import TXT to Text Items x64_7.dat
│   │   │   ├── HeDa_Import TXT to Text Items.lua
│   │   │   ├── Import Export TXT - changelog.txt
│   │   │   ├── Import Export TXT.dat
│   │   │   ├── License.txt
│   │   │   └── appicon.png
│   │   ├── Item Timecode
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Global Options.ini
│   │   │   ├── HITx32.dat
│   │   │   ├── HITx64.dat
│   │   │   ├── HRTx32.dat
│   │   │   ├── HRTx64.dat
│   │   │   ├── HeDa_Item Timecode.lua
│   │   │   ├── HeDa_Region Timecode.lua
│   │   │   ├── Item Timecode - changelog.txt
│   │   │   ├── Item Timecode.dat
│   │   │   └── License.txt
│   │   ├── Item Timecode 2
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Global Options.ini
│   │   │   ├── HITx32.dat
│   │   │   ├── HITx32_7.dat
│   │   │   ├── HITx64.dat
│   │   │   ├── HITx64_7.dat
│   │   │   ├── HRTx32.dat
│   │   │   ├── HRTx32_7.dat
│   │   │   ├── HRTx64.dat
│   │   │   ├── HRTx64_7.dat
│   │   │   ├── HeDa_Item Timecode 2.lua
│   │   │   ├── HeDa_Region Timecode 2.lua
│   │   │   ├── Item Timecode 2 - changelog.txt
│   │   │   ├── Item Timecode 2.dat
│   │   │   ├── License.txt
│   │   │   ├── appicon.png
│   │   │   └── background.png
│   │   ├── Item Timecode 2 settings
│   │   ├── Item Timecode settings
│   │   ├── MIDI Dynamics
│   │   │   ├── Default Settings.cfg
│   │   │   ├── HMDx32.dat
│   │   │   ├── HMDx32_7.dat
│   │   │   ├── HMDx64.dat
│   │   │   ├── HMDx64_7.dat
│   │   │   ├── HeDa_MIDI Dynamics.lua
│   │   │   ├── License.txt
│   │   │   ├── MIDI Dynamics - changelog.txt
│   │   │   ├── MIDI Dynamics.dat
│   │   │   └── appicon.png
│   │   ├── MIDI Dynamics settings
│   │   ├── Notes Reader
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Global Options.ini
│   │   │   ├── HNRx32.dat
│   │   │   ├── HNRx32_7.dat
│   │   │   ├── HNRx64.dat
│   │   │   ├── HNRx64_7.dat
│   │   │   ├── HeDa_Notes Reader.lua
│   │   │   ├── HeDa_Notes Reader_2nd.lua
│   │   │   ├── HeDa_Notes Reader_3rd.lua
│   │   │   ├── License.txt
│   │   │   ├── Notes Reader - changelog.txt
│   │   │   ├── Notes Reader.dat
│   │   │   └── appicon.png
│   │   ├── Notes Reader settings
│   │   ├── Projects
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Default themes
│   │   │   │   └── Default_6.0
│   │   │   │       ├── 200
│   │   │   │       │   ├── gen_play.png
│   │   │   │       │   ├── gen_play_on.png
│   │   │   │       │   └── gen_stop.png
│   │   │   │       ├── close.png
│   │   │   │       ├── gen_play.png
│   │   │   │       ├── gen_play_2.png
│   │   │   │       ├── gen_play_on.png
│   │   │   │       ├── gen_stop.png
│   │   │   │       ├── toolbar_card.png
│   │   │   │       ├── toolbar_grid.png
│   │   │   │       ├── toolbar_list.png
│   │   │   │       └── toolbar_menu.png
│   │   │   ├── HPx32.dat
│   │   │   ├── HPx32_7.dat
│   │   │   ├── HPx64.dat
│   │   │   ├── HPx64_7.dat
│   │   │   ├── HeDa Projects - Guide.pdf
│   │   │   ├── HeDa_Projects.lua
│   │   │   ├── License.txt
│   │   │   ├── Projects - changelog.txt
│   │   │   ├── Projects.dat
│   │   │   ├── appicon.png
│   │   │   ├── lib_HPx32.dat
│   │   │   ├── lib_HPx32_7.dat
│   │   │   ├── lib_HPx64.dat
│   │   │   └── lib_HPx64_7.dat
│   │   ├── Projects settings
│   │   │   └── audioplay
│   │   │       └── Projects_audioplay.RPP
│   │   ├── Regions&Markers from Items
│   │   │   ├── Global Options.ini
│   │   │   ├── HRFIx32.dat
│   │   │   ├── HRFIx64.dat
│   │   │   ├── HeDa_Regions&Markers from Items.lua
│   │   │   ├── HeDa_Regions&Markers run once without UI.lua
│   │   │   ├── Instructions.txt
│   │   │   ├── License.txt
│   │   │   ├── Regions&Markers from Items - changelog.txt
│   │   │   ├── Regions&Markers from Items.dat
│   │   │   └── appicon.png
│   │   ├── Remove Bypassed FX
│   │   │   ├── HRBFXx32.dat
│   │   │   ├── HRBFXx64.dat
│   │   │   ├── HeDa_Remove Bypassed FX.lua
│   │   │   ├── Remove Bypassed FX - changelog.txt
│   │   │   ├── Remove Bypassed FX.dat
│   │   │   └── appicon.png
│   │   ├── Render LUFS
│   │   │   ├── Default Settings.cfg
│   │   │   ├── HRLx32.dat
│   │   │   ├── HRLx64.dat
│   │   │   ├── HeDa_Render LUFS.lua
│   │   │   ├── License.txt
│   │   │   ├── Render LUFS - changelog.txt
│   │   │   ├── Render LUFS.dat
│   │   │   └── appicon.png
│   │   ├── Render LUFS settings
│   │   ├── Toggle Note Stretch to Fit
│   │   │   ├── HeDa_Toggle Note Stretch to Fit.eel
│   │   │   ├── Toggle Note Stretch to Fit - changelog.txt
│   │   │   ├── Toggle Note Stretch to Fit.dat
│   │   │   └── appicon.png
│   │   ├── Track Inspector
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Default themes
│   │   │   │   ├── Black
│   │   │   │   │   ├── mcp_env.png
│   │   │   │   │   ├── mcp_env_latch.png
│   │   │   │   │   ├── mcp_env_preview.png
│   │   │   │   │   ├── mcp_env_read.png
│   │   │   │   │   ├── mcp_env_touch.png
│   │   │   │   │   ├── mcp_env_write.png
│   │   │   │   │   ├── mcp_panthumb.png
│   │   │   │   │   ├── mcp_volthumb.png
│   │   │   │   │   └── meter.png
│   │   │   │   ├── Black.ti_theme
│   │   │   │   ├── Blue.ti_theme
│   │   │   │   ├── Default_3.0
│   │   │   │   │   ├── mcp_panthumb.png
│   │   │   │   │   ├── mcp_volthumb.png
│   │   │   │   │   ├── track_fx_dis.png
│   │   │   │   │   ├── track_fx_empty.png
│   │   │   │   │   ├── track_fx_norm.png
│   │   │   │   │   ├── track_fxempty_h.png
│   │   │   │   │   ├── track_fxoff_h.png
│   │   │   │   │   ├── track_fxon_h.png
│   │   │   │   │   ├── track_monitor_auto.png
│   │   │   │   │   ├── track_monitor_off.png
│   │   │   │   │   ├── track_monitor_on.png
│   │   │   │   │   ├── track_mute_off.png
│   │   │   │   │   ├── track_mute_on.png
│   │   │   │   │   ├── track_recarm_off.png
│   │   │   │   │   ├── track_recarm_on.png
│   │   │   │   │   ├── track_solo_off.png
│   │   │   │   │   ├── track_solo_on.png
│   │   │   │   │   ├── transport_record.png
│   │   │   │   │   └── transport_record_on.png
│   │   │   │   ├── Default_3.0.ti_theme
│   │   │   │   ├── Default_4.0
│   │   │   │   │   ├── mcp_env.png
│   │   │   │   │   ├── mcp_env_latch.png
│   │   │   │   │   ├── mcp_env_preview.png
│   │   │   │   │   ├── mcp_env_read.png
│   │   │   │   │   ├── mcp_env_touch.png
│   │   │   │   │   ├── mcp_env_write.png
│   │   │   │   │   ├── mcp_volthumb.png
│   │   │   │   │   ├── track_fx_empty.png
│   │   │   │   │   ├── track_fx_norm.png
│   │   │   │   │   ├── track_fxempty_h.png
│   │   │   │   │   ├── track_fxoff_h.png
│   │   │   │   │   ├── track_fxon_h.png
│   │   │   │   │   ├── track_monitor_auto.png
│   │   │   │   │   ├── track_monitor_off.png
│   │   │   │   │   ├── track_monitor_on.png
│   │   │   │   │   ├── track_mute_off.png
│   │   │   │   │   ├── track_mute_on.png
│   │   │   │   │   ├── track_recarm_off.png
│   │   │   │   │   ├── track_recarm_on.png
│   │   │   │   │   ├── track_solo_off.png
│   │   │   │   │   ├── track_solo_on.png
│   │   │   │   │   ├── transport_record.png
│   │   │   │   │   └── transport_record_on.png
│   │   │   │   ├── Default_4.0.ti_theme
│   │   │   │   ├── Default_5.0
│   │   │   │   │   ├── mcp_env.png
│   │   │   │   │   ├── mcp_env_latch.png
│   │   │   │   │   ├── mcp_env_preview.png
│   │   │   │   │   ├── mcp_env_read.png
│   │   │   │   │   ├── mcp_env_touch.png
│   │   │   │   │   ├── mcp_env_write.png
│   │   │   │   │   ├── mcp_panthumb.png
│   │   │   │   │   ├── mcp_volthumb.png
│   │   │   │   │   ├── meter.png
│   │   │   │   │   ├── meter_master.png
│   │   │   │   │   ├── toolbar_add_note.png
│   │   │   │   │   ├── toolbar_dock.png
│   │   │   │   │   ├── toolbar_freeze.png
│   │   │   │   │   ├── toolbar_freeze_version.png
│   │   │   │   │   ├── toolbar_unfreeze.png
│   │   │   │   │   ├── track_fx_empty.png
│   │   │   │   │   ├── track_fx_norm.png
│   │   │   │   │   ├── track_fxempty_h.png
│   │   │   │   │   ├── track_fxoff_h.png
│   │   │   │   │   ├── track_fxon_h.png
│   │   │   │   │   ├── track_monitor_auto.png
│   │   │   │   │   ├── track_monitor_off.png
│   │   │   │   │   ├── track_monitor_on.png
│   │   │   │   │   ├── track_mute_off.png
│   │   │   │   │   ├── track_mute_on.png
│   │   │   │   │   ├── track_recarm_off.png
│   │   │   │   │   ├── track_recarm_on.png
│   │   │   │   │   ├── track_solo_off.png
│   │   │   │   │   ├── track_solo_on.png
│   │   │   │   │   ├── transport_record_on_small.png
│   │   │   │   │   └── transport_record_small.png
│   │   │   │   ├── Default_5.0.ti_theme
│   │   │   │   ├── Default_5.0_@2x.ti_theme
│   │   │   │   ├── White
│   │   │   │   │   ├── toolbar_dock.png
│   │   │   │   │   ├── toolbar_freeze.png
│   │   │   │   │   ├── toolbar_freeze_version.png
│   │   │   │   │   └── toolbar_unfreeze.png
│   │   │   │   └── White.ti_theme
│   │   │   ├── FXGroups.cfg
│   │   │   ├── Global Options.ini
│   │   │   ├── HTIx32.dat
│   │   │   ├── HTIx64.dat
│   │   │   ├── HeDa_Track Inspector.lua
│   │   │   ├── HeDa_Track Inspector_Floating.lua
│   │   │   ├── HeDa_Track Inspector_Master.lua
│   │   │   ├── Installation Instructions.txt
│   │   │   ├── License.txt
│   │   │   ├── Track Inspector - changelog.txt
│   │   │   ├── Track Inspector.dat
│   │   │   └── appicon.png
│   │   ├── Track Inspector settings
│   │   ├── Track Templates
│   │   │   ├── Default Settings.cfg
│   │   │   ├── Default_5.0.TT_theme
│   │   │   ├── HTTx32.dat
│   │   │   ├── HTTx32_7.dat
│   │   │   ├── HTTx64.dat
│   │   │   ├── HTTx64_7.dat
│   │   │   ├── HeDa_Track Templates.lua
│   │   │   ├── License.txt
│   │   │   ├── Track Templates - changelog.txt
│   │   │   └── Track Templates.dat
│   │   └── Track Templates settings
│   ├── ReaTeam Extensions
│   │   └── API
│   │       ├── ReaImGui_Demo.lua
│   │       ├── gfx2imgui.lua
│   │       ├── imgui.lua
│   │       ├── imgui.py
│   │       └── sws.py
│   └── nvk-ReaScripts
│       ├── AUTODOPPLER
│       │   ├── Data
│       │   │   ├── custom.dat
│       │   │   ├── functions.dat
│       │   │   ├── grm.dat
│       │   │   ├── gui.dat
│       │   │   ├── nvk.dat
│       │   │   ├── scripts
│       │   │   │   └── autodoppler
│       │   │   │       ├── data
│       │   │   │       │   ├── actions.dat
│       │   │   │       │   ├── auto.dat
│       │   │   │       │   ├── bar.dat
│       │   │   │       │   ├── config.dat
│       │   │   │       │   ├── custom.dat
│       │   │   │       │   ├── dopplerpro.dat
│       │   │   │       │   ├── env.dat
│       │   │   │       │   ├── exit.dat
│       │   │   │       │   ├── frame.dat
│       │   │   │       │   ├── fx.dat
│       │   │   │       │   ├── grm.dat
│       │   │   │       │   ├── gui.dat
│       │   │   │       │   ├── loop.dat
│       │   │   │       │   ├── nvk.dat
│       │   │   │       │   ├── prefs.dat
│       │   │   │       │   ├── regions.dat
│       │   │   │       │   ├── render.dat
│       │   │   │       │   ├── soundparticles.dat
│       │   │   │       │   ├── spintracer.dat
│       │   │   │       │   ├── style.dat
│       │   │   │       │   ├── tabs.dat
│       │   │   │       │   ├── traveler.dat
│       │   │   │       │   ├── util.dat
│       │   │   │       │   └── waves.dat
│       │   │   │       └── main.dat
│       │   │   ├── soundparticles.dat
│       │   │   ├── traveler.dat
│       │   │   └── waves.dat
│       │   ├── Data53
│       │   │   ├── custom.dat
│       │   │   ├── functions.dat
│       │   │   ├── grm.dat
│       │   │   ├── gui.dat
│       │   │   ├── nvk.dat
│       │   │   ├── scripts
│       │   │   │   └── autodoppler
│       │   │   │       ├── data
│       │   │   │       │   ├── actions.dat
│       │   │   │       │   ├── auto.dat
│       │   │   │       │   ├── bar.dat
│       │   │   │       │   ├── config.dat
│       │   │   │       │   ├── custom.dat
│       │   │   │       │   ├── dopplerpro.dat
│       │   │   │       │   ├── env.dat
│       │   │   │       │   ├── exit.dat
│       │   │   │       │   ├── frame.dat
│       │   │   │       │   ├── fx.dat
│       │   │   │       │   ├── grm.dat
│       │   │   │       │   ├── gui.dat
│       │   │   │       │   ├── loop.dat
│       │   │   │       │   ├── nvk.dat
│       │   │   │       │   ├── prefs.dat
│       │   │   │       │   ├── regions.dat
│       │   │   │       │   ├── render.dat
│       │   │   │       │   ├── soundparticles.dat
│       │   │   │       │   ├── spintracer.dat
│       │   │   │       │   ├── style.dat
│       │   │   │       │   ├── tabs.dat
│       │   │   │       │   ├── traveler.dat
│       │   │   │       │   ├── util.dat
│       │   │   │       │   └── waves.dat
│       │   │   │       └── main.dat
│       │   │   ├── soundparticles.dat
│       │   │   ├── traveler.dat
│       │   │   └── waves.dat
│       │   ├── Presets
│       │   │   ├── TRAVELER.RPL
│       │   │   ├── TRAVELER_VST3.RPL
│       │   │   ├── vst-Traveler.ini
│       │   │   └── vst3-Traveler.ini
│       │   ├── nvk_AUTODOPPLER (legacy v1).lua
│       │   ├── nvk_AUTODOPPLER - Create snap offsets at rms peak.eel
│       │   ├── nvk_AUTODOPPLER - Custom.lua
│       │   ├── nvk_AUTODOPPLER - Quick Render - Randomize.lua
│       │   ├── nvk_AUTODOPPLER - Quick Render.lua
│       │   └── nvk_AUTODOPPLER.lua
│       ├── CREATE
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   ├── inv.cur
│       │   │   ├── replace.dat
│       │   │   └── scripts
│       │   │       └── create
│       │   │           ├── data
│       │   │           │   ├── db.dat
│       │   │           │   ├── gui.dat
│       │   │           │   ├── ml.dat
│       │   │           │   ├── mouse.dat
│       │   │           │   ├── peaksItem.dat
│       │   │           │   ├── resolution.dat
│       │   │           │   ├── results.dat
│       │   │           │   ├── retro.dat
│       │   │           │   ├── rpr.dat
│       │   │           │   ├── scr.dat
│       │   │           │   ├── search.dat
│       │   │           │   ├── soundminer.dat
│       │   │           │   ├── string.dat
│       │   │           │   ├── tbl.dat
│       │   │           │   ├── timers.dat
│       │   │           │   └── track.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   ├── replace.dat
│       │   │   └── scripts
│       │   │       └── create
│       │   │           ├── data
│       │   │           │   ├── db.dat
│       │   │           │   ├── gui.dat
│       │   │           │   ├── ml.dat
│       │   │           │   ├── mouse.dat
│       │   │           │   ├── peaksItem.dat
│       │   │           │   ├── resolution.dat
│       │   │           │   ├── results.dat
│       │   │           │   ├── retro.dat
│       │   │           │   ├── rpr.dat
│       │   │           │   ├── scr.dat
│       │   │           │   ├── search.dat
│       │   │           │   ├── soundminer.dat
│       │   │           │   ├── string.dat
│       │   │           │   ├── tbl.dat
│       │   │           │   ├── timers.dat
│       │   │           │   └── track.dat
│       │   │           └── main.dat
│       │   ├── nvk_CREATE - Replace.lua
│       │   ├── nvk_CREATE.lua
│       │   ├── nvk_CREATE_cfg
│       │   ├── nvk_CREATE_fav
│       │   └── nvk_CREATE_hide
│       ├── FOLDER_ITEMS
│       │   ├── Data
│       │   │   ├── folder_items.dat
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       ├── rename
│       │   │       │   ├── data
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── advanced.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── categories.dat
│       │   │       │   │   ├── cb.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── exit.dat
│       │   │       │   │   ├── file_list.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── gui.dat
│       │   │       │   │   ├── input.dat
│       │   │       │   │   ├── keyboard.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── options.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   ├── targets.dat
│       │   │       │   │   ├── tt.dat
│       │   │       │   │   └── ucs.dat
│       │   │       │   └── main.dat
│       │   │       ├── render_smart
│       │   │       │   ├── data
│       │   │       │   │   ├── action.dat
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── cb.dat
│       │   │       │   │   ├── colors.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── exit.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── gui.dat
│       │   │       │   │   ├── keyboard.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   ├── render_list.dat
│       │   │       │   │   ├── render_settings.dat
│       │   │       │   │   ├── sausage.dat
│       │   │       │   │   ├── scr.dat
│       │   │       │   │   ├── tabs.dat
│       │   │       │   │   ├── tt.dat
│       │   │       │   │   └── util.dat
│       │   │       │   └── main.dat
│       │   │       ├── reposition
│       │   │       │   ├── data
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── cb.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── exit.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── gui.dat
│       │   │       │   │   ├── list.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   ├── style.dat
│       │   │       │   │   ├── tabs.dat
│       │   │       │   │   └── tt.dat
│       │   │       │   └── main.dat
│       │   │       └── settings
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── cb.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── exit.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── gui.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── mm.dat
│       │   │           │   ├── tabs.dat
│       │   │           │   └── tt.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── folder_items.dat
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       ├── rename
│       │   │       │   ├── data
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── advanced.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── categories.dat
│       │   │       │   │   ├── cb.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── exit.dat
│       │   │       │   │   ├── file_list.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── gui.dat
│       │   │       │   │   ├── input.dat
│       │   │       │   │   ├── keyboard.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── options.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   ├── targets.dat
│       │   │       │   │   ├── tt.dat
│       │   │       │   │   └── ucs.dat
│       │   │       │   └── main.dat
│       │   │       ├── render_smart
│       │   │       │   ├── data
│       │   │       │   │   ├── action.dat
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── cb.dat
│       │   │       │   │   ├── colors.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── exit.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── gui.dat
│       │   │       │   │   ├── keyboard.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   ├── render_list.dat
│       │   │       │   │   ├── render_settings.dat
│       │   │       │   │   ├── sausage.dat
│       │   │       │   │   ├── scr.dat
│       │   │       │   │   ├── tabs.dat
│       │   │       │   │   ├── tt.dat
│       │   │       │   │   └── util.dat
│       │   │       │   └── main.dat
│       │   │       ├── reposition
│       │   │       │   ├── data
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── cb.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── exit.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── gui.dat
│       │   │       │   │   ├── list.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   ├── style.dat
│       │   │       │   │   ├── tabs.dat
│       │   │       │   │   └── tt.dat
│       │   │       │   └── main.dat
│       │   │       └── settings
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── cb.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── exit.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── gui.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── mm.dat
│       │   │           │   ├── tabs.dat
│       │   │           │   └── tt.dat
│       │   │           └── main.dat
│       │   ├── nvk_FOLDER_ITEMS - Add new items to existing folder.lua
│       │   ├── nvk_FOLDER_ITEMS - Copy.lua
│       │   ├── nvk_FOLDER_ITEMS - Create new folder from selected items or tracks.lua
│       │   ├── nvk_FOLDER_ITEMS - Cut.lua
│       │   ├── nvk_FOLDER_ITEMS - Deselect folder items.lua
│       │   ├── nvk_FOLDER_ITEMS - Deselect non-folder items.lua
│       │   ├── nvk_FOLDER_ITEMS - Deselect non-render items SMART.lua
│       │   ├── nvk_FOLDER_ITEMS - Fade SMART.lua
│       │   ├── nvk_FOLDER_ITEMS - Fade in.lua
│       │   ├── nvk_FOLDER_ITEMS - Fade out.lua
│       │   ├── nvk_FOLDER_ITEMS - Item Render Settings - Channels - Custom.lua
│       │   ├── nvk_FOLDER_ITEMS - Item Render Settings - Channels - Mono.lua
│       │   ├── nvk_FOLDER_ITEMS - Item Render Settings - Channels - Stereo.lua
│       │   ├── nvk_FOLDER_ITEMS - Join.lua
│       │   ├── nvk_FOLDER_ITEMS - MM - Media Item - Double Click.lua
│       │   ├── nvk_FOLDER_ITEMS - MM - Reposition - Groups.lua
│       │   ├── nvk_FOLDER_ITEMS - MM - Reposition - Items.lua
│       │   ├── nvk_FOLDER_ITEMS - MM - Reposition - Propagate.lua
│       │   ├── nvk_FOLDER_ITEMS - MM - TCP - Double Click.lua
│       │   ├── nvk_FOLDER_ITEMS - MM - Track - Double Click.lua
│       │   ├── nvk_FOLDER_ITEMS - Mousewheel Pitch Shift - Octave (Reversed).lua
│       │   ├── nvk_FOLDER_ITEMS - Mousewheel Pitch Shift - Octave.lua
│       │   ├── nvk_FOLDER_ITEMS - Mousewheel Pitch Shift - Semitone (Reversed).lua
│       │   ├── nvk_FOLDER_ITEMS - Mousewheel Pitch Shift - Semitone.lua
│       │   ├── nvk_FOLDER_ITEMS - Mousewheel Volume.lua
│       │   ├── nvk_FOLDER_ITEMS - Organize - Cleanup RENDERS folder.lua
│       │   ├── nvk_FOLDER_ITEMS - Organize - Items to as few tracks as possible (preserve column order).lua
│       │   ├── nvk_FOLDER_ITEMS - Organize - Items to as few tracks as possible.lua
│       │   ├── nvk_FOLDER_ITEMS - Organize - Keep tracks with items in order by time.lua
│       │   ├── nvk_FOLDER_ITEMS - Paste.lua
│       │   ├── nvk_FOLDER_ITEMS - Remove.lua
│       │   ├── nvk_FOLDER_ITEMS - Rename.lua
│       │   ├── nvk_FOLDER_ITEMS - Render SMART.lua
│       │   ├── nvk_FOLDER_ITEMS - Render.lua
│       │   ├── nvk_FOLDER_ITEMS - Reposition - Preset 1.lua
│       │   ├── nvk_FOLDER_ITEMS - Reposition - Preset 2.lua
│       │   ├── nvk_FOLDER_ITEMS - Reposition - Preset 3.lua
│       │   ├── nvk_FOLDER_ITEMS - Reposition - Preset 4.lua
│       │   ├── nvk_FOLDER_ITEMS - Reposition.lua
│       │   ├── nvk_FOLDER_ITEMS - Settings.lua
│       │   ├── nvk_FOLDER_ITEMS - Split.lua
│       │   ├── nvk_FOLDER_ITEMS - Trim left edge.lua
│       │   ├── nvk_FOLDER_ITEMS - Trim right edge.lua
│       │   ├── nvk_FOLDER_ITEMS - Update (manual).lua
│       │   └── nvk_FOLDER_ITEMS.lua
│       ├── FONTS
│       │   └── -ADD-FONTS-TO-THIS-FOLDER-
│       ├── ICONS
│       │   └── nvk_ICONS.lua
│       ├── ITEMS
│       │   ├── Data
│       │   │   └── functions.dat
│       │   ├── Data53
│       │   │   └── functions.dat
│       │   ├── nvk_ITEMS - Align SMART.lua
│       │   ├── nvk_ITEMS - Apply FX to items and move to directory.lua
│       │   ├── nvk_ITEMS - Copy item positions.lua
│       │   ├── nvk_ITEMS - Copy item volumes.lua
│       │   ├── nvk_ITEMS - Fade curve SMART (mousewheel reversed).lua
│       │   ├── nvk_ITEMS - Fade curve SMART (mousewheel).lua
│       │   ├── nvk_ITEMS - Fade in curve (mousewheel reversed) .lua
│       │   ├── nvk_ITEMS - Fade in curve (mousewheel).lua
│       │   ├── nvk_ITEMS - Fade out curve (mousewheel reversed).lua
│       │   ├── nvk_ITEMS - Fade out curve (mousewheel).lua
│       │   ├── nvk_ITEMS - Fades to volume automation.lua
│       │   ├── nvk_ITEMS - Glue all selected items independently and set source name to active take name (add numbers if file already exists).lua
│       │   ├── nvk_ITEMS - Glue all selected items independently and set source name to active take name.lua
│       │   ├── nvk_ITEMS - Halve item volume changes.lua
│       │   ├── nvk_ITEMS - Move cursor to next transient.lua
│       │   ├── nvk_ITEMS - Move cursor to previous transient.lua
│       │   ├── nvk_ITEMS - Move selected items down one track SMART.lua
│       │   ├── nvk_ITEMS - Move selected items up one track SMART.lua
│       │   ├── nvk_ITEMS - Move to new tracks with track fx.lua
│       │   ├── nvk_ITEMS - Paste item positions.lua
│       │   ├── nvk_ITEMS - Paste item volumes (reverse gain).lua
│       │   ├── nvk_ITEMS - Paste item volumes.lua
│       │   ├── nvk_ITEMS - Project marker to item notes.lua
│       │   ├── nvk_ITEMS - Remove muted items from selection.lua
│       │   ├── nvk_ITEMS - Remove selected items greater or less than length.lua
│       │   ├── nvk_ITEMS - Remove selected items with no color.lua
│       │   ├── nvk_ITEMS - Reposition selected items (ignore snap offset).lua
│       │   ├── nvk_ITEMS - Reposition selected items across tracks.lua
│       │   ├── nvk_ITEMS - Reposition selected items.lua
│       │   ├── nvk_ITEMS - Reset item length to take length - remove fades - reset volume.lua
│       │   ├── nvk_ITEMS - Reset item length to take length.lua
│       │   ├── nvk_ITEMS - Reset item to original state.lua
│       │   ├── nvk_ITEMS - Reset snap offsets of selected items.lua
│       │   ├── nvk_ITEMS - Select all items after selected items.lua
│       │   ├── nvk_ITEMS - Select all items before selected items.lua
│       │   ├── nvk_ITEMS - Select and move edit cursor to next item and preserve play state.lua
│       │   ├── nvk_ITEMS - Select and move edit cursor to next item edge.lua
│       │   ├── nvk_ITEMS - Select and move edit cursor to previous item and preserve play state.lua
│       │   ├── nvk_ITEMS - Select and move edit cursor to previous item edge.lua
│       │   ├── nvk_ITEMS - Shift take offset by length of item.lua
│       │   ├── nvk_ITEMS - Sort groups of variations of selected items to new tracks based on name and name tracks with folder structure.lua
│       │   ├── nvk_ITEMS - Sort groups of variations of selected items to new tracks based on name and name tracks.lua
│       │   ├── nvk_ITEMS - Sort groups of variations of selected items to new tracks based on name.lua
│       │   ├── nvk_ITEMS - Split selected items x length x space.lua
│       │   ├── nvk_ITEMS - Split selected items x number of times.lua
│       │   ├── nvk_ITEMS - Stretch selected items to previous and next markers.lua
│       │   ├── nvk_ITEMS - Swap selected items tracks.lua
│       │   ├── nvk_ITEMS - Switch item source to next in folder.lua
│       │   ├── nvk_ITEMS - Switch item source to previous in folder.lua
│       │   ├── nvk_ITEMS - Toggle mute SMART.lua
│       │   ├── nvk_ITEMS - Trim overlapping items to smallest item in each column.lua
│       │   ├── nvk_ITEMS - Unselect muted items.lua
│       │   └── nvk_ITEMS.lua
│       ├── LOOPMAKER
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   ├── loopmaker.dat
│       │   │   └── scripts
│       │   │       └── loopmaker
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── bar.dat
│       │   │           │   ├── cb.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── exit.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── list.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── name.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   ├── loopmaker.dat
│       │   │   └── scripts
│       │   │       └── loopmaker
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── bar.dat
│       │   │           │   ├── cb.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── exit.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── list.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── name.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── nvk_LOOPMAKER (legacy v1).lua
│       │   └── nvk_LOOPMAKER.lua
│       ├── MULTI
│       │   ├── nvk_MULTI - Copy item positions and item volumes and take name.lua
│       │   ├── nvk_MULTI - Copy item positions and take name.lua
│       │   ├── nvk_MULTI - Insert or edit marker at current position.lua
│       │   ├── nvk_MULTI - Move tracks-items-envelope points down depending on focus SMART.lua
│       │   ├── nvk_MULTI - Move tracks-items-envelope points up depending on focus SMART.lua
│       │   ├── nvk_MULTI - Move tracks-items-envelope points up or down depending on focus SMART (mousewheel reversed).lua
│       │   ├── nvk_MULTI - Move tracks-items-envelope points up or down depending on focus SMART (mousewheel).lua
│       │   ├── nvk_MULTI - Project organize - cleanup - color (multitap).lua
│       │   ├── nvk_MULTI - Project organize - cleanup - color special (multitap).lua
│       │   ├── nvk_MULTI - Remove bypassed and offline FX from selected tracks or muted items from selection depending on focus.lua
│       │   ├── nvk_MULTI - Remove time selection and loop points and deselect all items.lua
│       │   ├── nvk_MULTI - Toggle fx bypass depending on focus.lua
│       │   └── nvk_MULTI.lua
│       ├── PROPAGATE
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── propagate
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── bar.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── propagate
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── bar.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── nvk_PROPAGATE - Apply.lua
│       │   └── nvk_PROPAGATE.lua
│       ├── SEARCH
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── search
│       │   │           ├── data
│       │   │           │   ├── action.dat
│       │   │           │   ├── actions.dat
│       │   │           │   ├── bar.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── dnd.dat
│       │   │           │   ├── exit.dat
│       │   │           │   ├── folder.dat
│       │   │           │   ├── folders.dat
│       │   │           │   ├── font.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── fx.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── markers.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   ├── project.dat
│       │   │           │   ├── recent.dat
│       │   │           │   ├── regions.dat
│       │   │           │   ├── result.dat
│       │   │           │   ├── results.dat
│       │   │           │   ├── search.dat
│       │   │           │   ├── takes.dat
│       │   │           │   └── tracks.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── search
│       │   │           ├── data
│       │   │           │   ├── action.dat
│       │   │           │   ├── actions.dat
│       │   │           │   ├── bar.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── dnd.dat
│       │   │           │   ├── exit.dat
│       │   │           │   ├── folder.dat
│       │   │           │   ├── folders.dat
│       │   │           │   ├── font.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── fx.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── markers.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   ├── project.dat
│       │   │           │   ├── recent.dat
│       │   │           │   ├── regions.dat
│       │   │           │   ├── result.dat
│       │   │           │   ├── results.dat
│       │   │           │   ├── search.dat
│       │   │           │   ├── takes.dat
│       │   │           │   └── tracks.dat
│       │   │           └── main.dat
│       │   ├── nvk_SEARCH - Action.lua
│       │   ├── nvk_SEARCH - All.lua
│       │   ├── nvk_SEARCH - Chain.lua
│       │   ├── nvk_SEARCH - FX.lua
│       │   ├── nvk_SEARCH - Folder Item.lua
│       │   ├── nvk_SEARCH - Marker.lua
│       │   ├── nvk_SEARCH - Project.lua
│       │   ├── nvk_SEARCH - Region.lua
│       │   ├── nvk_SEARCH - Take.lua
│       │   ├── nvk_SEARCH - Track Template.lua
│       │   ├── nvk_SEARCH - Track.lua
│       │   ├── nvk_SEARCH.lua
│       │   └── nvk_SEARCH_cfg
│       ├── SHARED
│       │   ├── Data
│       │   │   ├── basics.dat
│       │   │   ├── data
│       │   │   │   ├── actions.dat
│       │   │   │   ├── bar.dat
│       │   │   │   ├── cb.dat
│       │   │   │   ├── colors.dat
│       │   │   │   ├── demo.dat
│       │   │   │   ├── exit.dat
│       │   │   │   ├── file.dat
│       │   │   │   ├── font.dat
│       │   │   │   ├── fx.dat
│       │   │   │   ├── gui.dat
│       │   │   │   ├── key.dat
│       │   │   │   ├── keyboard.dat
│       │   │   │   ├── loop.dat
│       │   │   │   ├── mouse.dat
│       │   │   │   ├── prefs.dat
│       │   │   │   ├── prof.dat
│       │   │   │   ├── result.dat
│       │   │   │   ├── stats.dat
│       │   │   │   ├── style.dat
│       │   │   │   ├── tabs.dat
│       │   │   │   ├── terminal.dat
│       │   │   │   ├── text.dat
│       │   │   │   ├── theme.dat
│       │   │   │   ├── tt.dat
│       │   │   │   └── upgrade.dat
│       │   │   ├── fonts
│       │   │   │   └── icons.otf
│       │   │   ├── functions.dat
│       │   │   ├── license.dat
│       │   │   ├── load.dat
│       │   │   ├── main.dat
│       │   │   ├── meta
│       │   │   │   ├── autoitem.dat
│       │   │   │   ├── autoitems.dat
│       │   │   │   ├── class.dat
│       │   │   │   ├── clip.dat
│       │   │   │   ├── clips.dat
│       │   │   │   ├── col.dat
│       │   │   │   ├── column.dat
│       │   │   │   ├── columns.dat
│       │   │   │   ├── envelope.dat
│       │   │   │   ├── folderitem.dat
│       │   │   │   ├── folderitems.dat
│       │   │   │   ├── item.dat
│       │   │   │   ├── items.dat
│       │   │   │   ├── marker.dat
│       │   │   │   ├── markers.dat
│       │   │   │   ├── take.dat
│       │   │   │   ├── takemarker.dat
│       │   │   │   ├── takemarkers.dat
│       │   │   │   ├── takes.dat
│       │   │   │   ├── track.dat
│       │   │   │   └── tracks.dat
│       │   │   ├── meta.dat
│       │   │   ├── scr.dat
│       │   │   ├── scripts
│       │   │   │   └── manager
│       │   │   │       ├── data
│       │   │   │       │   ├── actions.dat
│       │   │   │       │   ├── config.dat
│       │   │   │       │   ├── frame.dat
│       │   │   │       │   ├── loop.dat
│       │   │   │       │   ├── prefs.dat
│       │   │   │       │   └── tabs.dat
│       │   │   │       └── main.dat
│       │   │   ├── simple
│       │   │   │   ├── data
│       │   │   │   │   ├── actions.dat
│       │   │   │   │   ├── bar.dat
│       │   │   │   │   ├── config.dat
│       │   │   │   │   ├── frame.dat
│       │   │   │   │   ├── loop.dat
│       │   │   │   │   ├── prefs.dat
│       │   │   │   │   └── tabs.dat
│       │   │   │   └── main.dat
│       │   │   └── template
│       │   │       ├── data
│       │   │       │   ├── actions.dat
│       │   │       │   ├── config.dat
│       │   │       │   ├── frame.dat
│       │   │       │   ├── loop.dat
│       │   │       │   ├── prefs.dat
│       │   │       │   └── tabs.dat
│       │   │       ├── functions.dat
│       │   │       └── main.dat
│       │   ├── Data53
│       │   │   ├── basics.dat
│       │   │   ├── data
│       │   │   │   ├── actions.dat
│       │   │   │   ├── bar.dat
│       │   │   │   ├── cb.dat
│       │   │   │   ├── colors.dat
│       │   │   │   ├── demo.dat
│       │   │   │   ├── exit.dat
│       │   │   │   ├── file.dat
│       │   │   │   ├── font.dat
│       │   │   │   ├── fx.dat
│       │   │   │   ├── gui.dat
│       │   │   │   ├── key.dat
│       │   │   │   ├── keyboard.dat
│       │   │   │   ├── loop.dat
│       │   │   │   ├── mouse.dat
│       │   │   │   ├── prefs.dat
│       │   │   │   ├── prof.dat
│       │   │   │   ├── result.dat
│       │   │   │   ├── stats.dat
│       │   │   │   ├── style.dat
│       │   │   │   ├── tabs.dat
│       │   │   │   ├── terminal.dat
│       │   │   │   ├── text.dat
│       │   │   │   ├── theme.dat
│       │   │   │   ├── tt.dat
│       │   │   │   └── upgrade.dat
│       │   │   ├── functions.dat
│       │   │   ├── license.dat
│       │   │   ├── load.dat
│       │   │   ├── main.dat
│       │   │   ├── meta
│       │   │   │   ├── autoitem.dat
│       │   │   │   ├── autoitems.dat
│       │   │   │   ├── class.dat
│       │   │   │   ├── clip.dat
│       │   │   │   ├── clips.dat
│       │   │   │   ├── col.dat
│       │   │   │   ├── column.dat
│       │   │   │   ├── columns.dat
│       │   │   │   ├── envelope.dat
│       │   │   │   ├── folderitem.dat
│       │   │   │   ├── folderitems.dat
│       │   │   │   ├── item.dat
│       │   │   │   ├── items.dat
│       │   │   │   ├── marker.dat
│       │   │   │   ├── markers.dat
│       │   │   │   ├── take.dat
│       │   │   │   ├── takemarker.dat
│       │   │   │   ├── takemarkers.dat
│       │   │   │   ├── takes.dat
│       │   │   │   ├── track.dat
│       │   │   │   └── tracks.dat
│       │   │   ├── meta.dat
│       │   │   ├── scr.dat
│       │   │   ├── scripts
│       │   │   │   └── manager
│       │   │   │       ├── data
│       │   │   │       │   ├── actions.dat
│       │   │   │       │   ├── config.dat
│       │   │   │       │   ├── frame.dat
│       │   │   │       │   ├── loop.dat
│       │   │   │       │   ├── prefs.dat
│       │   │   │       │   └── tabs.dat
│       │   │   │       └── main.dat
│       │   │   ├── simple
│       │   │   │   ├── data
│       │   │   │   │   ├── actions.dat
│       │   │   │   │   ├── bar.dat
│       │   │   │   │   ├── config.dat
│       │   │   │   │   ├── frame.dat
│       │   │   │   │   ├── loop.dat
│       │   │   │   │   ├── prefs.dat
│       │   │   │   │   └── tabs.dat
│       │   │   │   └── main.dat
│       │   │   └── template
│       │   │       ├── data
│       │   │       │   ├── actions.dat
│       │   │       │   ├── config.dat
│       │   │       │   ├── frame.dat
│       │   │       │   ├── loop.dat
│       │   │       │   ├── prefs.dat
│       │   │       │   └── tabs.dat
│       │   │       ├── functions.dat
│       │   │       └── main.dat
│       │   ├── nvk_SHARED - License Manager.lua
│       │   └── nvk_SHARED.lua
│       ├── SUBPROJECT
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       ├── settings
│       │   │       │   ├── data
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   └── tabs.dat
│       │   │       │   └── main.dat
│       │   │       └── subproject
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── fx.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   ├── result.dat
│       │   │           │   ├── results.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       ├── settings
│       │   │       │   ├── data
│       │   │       │   │   ├── actions.dat
│       │   │       │   │   ├── bar.dat
│       │   │       │   │   ├── config.dat
│       │   │       │   │   ├── frame.dat
│       │   │       │   │   ├── loop.dat
│       │   │       │   │   ├── prefs.dat
│       │   │       │   │   └── tabs.dat
│       │   │       │   └── main.dat
│       │   │       └── subproject
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── fx.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   ├── result.dat
│       │   │           │   ├── results.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── nvk_SUBPROJECT - Fix markers.lua
│       │   ├── nvk_SUBPROJECT - Settings.lua
│       │   ├── nvk_SUBPROJECT - Update - Match main project positions.lua
│       │   ├── nvk_SUBPROJECT - Update - Match subproject positions.lua
│       │   ├── nvk_SUBPROJECT - Update - Relative positions.lua
│       │   ├── nvk_SUBPROJECT - Update - Render (no reposition).lua
│       │   └── nvk_SUBPROJECT.lua
│       ├── TAKES
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── settings
│       │   │           ├── data
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── loop.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── settings
│       │   │           ├── data
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── loop.dat
│       │   │           │   └── tabs.dat
│       │   │           └── main.dat
│       │   ├── nvk_TAKES - Add named take marker to start of selected items (first track only).lua
│       │   ├── nvk_TAKES - Add named take marker to start of selected items.lua
│       │   ├── nvk_TAKES - Add stretch markers at start and end of selected items takes and set slope.lua
│       │   ├── nvk_TAKES - Add stretch markers at start and end of selected items takes.lua
│       │   ├── nvk_TAKES - Add take markers to all variations in selected items takes.lua
│       │   ├── nvk_TAKES - Add take markers to variations in items.eel
│       │   ├── nvk_TAKES - Consolidate item splits as take markers in first item.lua
│       │   ├── nvk_TAKES - Consolidate takes with take markers SMART.eel
│       │   ├── nvk_TAKES - Convert 5.1 to Quad.lua
│       │   ├── nvk_TAKES - Copy take name(s).lua
│       │   ├── nvk_TAKES - Copy take name.lua
│       │   ├── nvk_TAKES - Duplicate items and select next take SMART.lua
│       │   ├── nvk_TAKES - Find takes by name.lua
│       │   ├── nvk_TAKES - Paste copied take names to selected items.lua
│       │   ├── nvk_TAKES - Quick add numbered take marker at mouse position.lua
│       │   ├── nvk_TAKES - Render to new take with fx.lua
│       │   ├── nvk_TAKES - Reverse.lua
│       │   ├── nvk_TAKES - Select next take SMART.lua
│       │   ├── nvk_TAKES - Select next take.lua
│       │   ├── nvk_TAKES - Select previous take SMART.lua
│       │   ├── nvk_TAKES - Select previous take.lua
│       │   ├── nvk_TAKES - Select random take SMART (no repeat).lua
│       │   ├── nvk_TAKES - Select random take SMART.lua
│       │   ├── nvk_TAKES - Settings.lua
│       │   ├── nvk_TAKES - Toggle width fx - save last touched parameter as script setting.lua
│       │   ├── nvk_TAKES - Toggle width fx or toggle track width envelope.lua
│       │   └── nvk_TAKES.lua
│       ├── THEME
│       │   ├── Data
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── settings
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   ├── tabs.dat
│       │   │           │   └── theme.dat
│       │   │           └── main.dat
│       │   ├── Data53
│       │   │   ├── functions.dat
│       │   │   └── scripts
│       │   │       └── settings
│       │   │           ├── data
│       │   │           │   ├── actions.dat
│       │   │           │   ├── config.dat
│       │   │           │   ├── frame.dat
│       │   │           │   ├── loop.dat
│       │   │           │   ├── prefs.dat
│       │   │           │   ├── tabs.dat
│       │   │           │   └── theme.dat
│       │   │           └── main.dat
│       │   ├── nvk_THEME - Output - Quad.lua
│       │   ├── nvk_THEME - Output - Stereo.lua
│       │   ├── nvk_THEME - Output - Surround.lua
│       │   ├── nvk_THEME - Settings.lua
│       │   └── nvk_THEME.lua
│       ├── TRACK
│       │   ├── Data
│       │   │   └── functions.dat
│       │   ├── Data53
│       │   │   └── functions.dat
│       │   ├── nvk_TRACK - Close all track FX windows.lua
│       │   ├── nvk_TRACK - Color folder and named tracks with SWS custom colors (parent tracks only but color all items).lua
│       │   ├── nvk_TRACK - Color folder and named tracks with SWS custom colors (parent tracks only).lua
│       │   ├── nvk_TRACK - Color folder and named tracks with SWS custom colors.lua
│       │   ├── nvk_TRACK - Cycle folder state of all tracks.lua
│       │   ├── nvk_TRACK - Duplicate tracks or items depending on focus.lua
│       │   ├── nvk_TRACK - Insert and set color to parent track color.lua
│       │   ├── nvk_TRACK - Move folder and named tracks to bottom of project and video track to top.lua
│       │   ├── nvk_TRACK - Move folder and named tracks to top of project and video track to top.lua
│       │   ├── nvk_TRACK - Move selected tracks down SMART.lua
│       │   ├── nvk_TRACK - Move selected tracks up SMART.lua
│       │   ├── nvk_TRACK - Remove bypassed and offline FX from selected tracks.lua
│       │   ├── nvk_TRACK - Remove unused tracks.lua
│       │   ├── nvk_TRACK - Reset track or item color depending on focus.lua
│       │   ├── nvk_TRACK - Set master track channel count.lua
│       │   ├── nvk_TRACK - Set selected tracks channel count.lua
│       │   ├── nvk_TRACK - Toggle mute tracks of selected items.lua
│       │   ├── nvk_TRACK - Toggle solo tracks of selected items - mute other tracks.lua
│       │   ├── nvk_TRACK - Toggle solo tracks of selected items.lua
│       │   └── nvk_TRACK.lua
│       └── VARIATIONS
│           ├── Data
│           │   ├── functions.dat
│           │   └── scripts
│           │       └── variations
│           │           ├── data
│           │           │   ├── actions.dat
│           │           │   ├── bar.dat
│           │           │   ├── cb.dat
│           │           │   ├── config.dat
│           │           │   ├── exit.dat
│           │           │   ├── frame.dat
│           │           │   ├── gui.dat
│           │           │   ├── list.dat
│           │           │   ├── loop.dat
│           │           │   ├── prefs.dat
│           │           │   ├── tabs.dat
│           │           │   ├── tt.dat
│           │           │   ├── undo.dat
│           │           │   └── var.dat
│           │           └── main.dat
│           ├── Data53
│           │   ├── functions.dat
│           │   └── scripts
│           │       └── variations
│           │           ├── data
│           │           │   ├── actions.dat
│           │           │   ├── bar.dat
│           │           │   ├── cb.dat
│           │           │   ├── config.dat
│           │           │   ├── exit.dat
│           │           │   ├── frame.dat
│           │           │   ├── gui.dat
│           │           │   ├── list.dat
│           │           │   ├── loop.dat
│           │           │   ├── prefs.dat
│           │           │   ├── tabs.dat
│           │           │   ├── tt.dat
│           │           │   ├── undo.dat
│           │           │   └── var.dat
│           │           └── main.dat
│           ├── nvk_VARIATIONS - New Variation - Preset 0 (default).lua
│           ├── nvk_VARIATIONS - New Variation - Preset 1.lua
│           ├── nvk_VARIATIONS - New Variation - Preset 2.lua
│           ├── nvk_VARIATIONS - New Variation - Preset 3.lua
│           ├── nvk_VARIATIONS - New Variation - Preset 4.lua
│           ├── nvk_VARIATIONS - Randomize - Preset 0 (default).lua
│           ├── nvk_VARIATIONS - Randomize - Preset 1.lua
│           ├── nvk_VARIATIONS - Randomize - Preset 2.lua
│           ├── nvk_VARIATIONS - Randomize - Preset 3.lua
│           ├── nvk_VARIATIONS - Randomize - Preset 4.lua
│           └── nvk_VARIATIONS.lua
├── TrackTemplates
├── UserPlugins
│   ├── reaper_imgui-arm64.dylib
│   ├── reaper_js_ReaScriptAPI64ARM.dylib
│   ├── reaper_reapack-arm64.dylib
│   └── reaper_sws-arm64.dylib
├── presets
│   └── vst3-Kontakt 7-builtin.ini
├── reapack.ini
├── reaper-auplugins_arm64-bc.ini
├── reaper-auplugins_arm64.ini
├── reaper-clap-macos-aarch64.ini
├── reaper-extstate.ini
├── reaper-fxtags.ini
├── reaper-install-rev.txt
├── reaper-jsfx.ini
├── reaper-kb.ini
├── reaper-license.rk
├── reaper-midihw.ini
├── reaper-mouse.ini
├── reaper-recentfx.ini
├── reaper-reginfo2.ini
├── reaper-themeconfig.ini
├── reaper-vstplugins_arm64.ini
├── reaper-vstshells_arm64.ini
├── reaper-wndpos.ini
├── reaper.ini
└── sws-autocoloricon.ini
```

package models

type SavedSearchesResponse struct {
	Entry    []SavedSearchesEntry `json:"entry"`
	Messages []ErrorMessage       `json:"messages"`
}

type SavedSearchesEntry struct {
	Name    string            `json:"name"`
	ACL     ACLObject         `json:"acl"`
	Content SavedSearchObject `json:"content"`
}

type SavedSearchObject struct {
	Actions                                   string  `json:"actions,omitempty" url:"actions,omitempty"`
	ActionEmail                               bool    `json:"action.email,omitempty" url:"action.email"`
	ActionEmailAuthPassword                   string  `json:"action.email.auth_password,omitempty" url:"action.email.auth_password,omitempty"`
	ActionEmailAuthUsername                   string  `json:"action.email.auth_username,omitempty" url:"action.email.auth_username,omitempty"`
	ActionEmailBCC                            string  `json:"action.email.bcc,omitempty" url:"action.email.bcc,omitempty"`
	ActionEmailCC                             string  `json:"action.email.cc,omitempty" url:"action.email.cc,omitempty"`
	ActionEmailCommand                        string  `json:"action.email.command,omitempty" url:"action.email.command,omitempty"`
	ActionEmailFormat                         string  `json:"action.email.format,omitempty" url:"action.email.format,omitempty"`
	ActionEmailFrom                           string  `json:"action.email.from,omitempty" url:"action.email.from,omitempty"`
	ActionEmailHostname                       string  `json:"action.email.hostname,omitempty" url:"action.email.hostname,omitempty"`
	ActionEmailIncludeResultsLink             int     `json:"action.email.include.results_link,string,omitempty" url:"action.email.include.results_link,omitempty"`
	ActionEmailIncludeSearch                  int     `json:"action.email.include.search,string,omitempty" url:"action.email.include.search,omitempty"`
	ActionEmailIncludeTrigger                 int     `json:"action.email.include.trigger,string,omitempty" url:"action.email.include.trigger,omitempty"`
	ActionEmailIncludeTriggerTime             int     `json:"action.email.include.trigger_time,string,omitempty" url:"action.email.include.trigger_time,omitempty"`
	ActionEmailIncludeViewLink                int     `json:"action.email.include.view_link,string,omitempty" url:"action.email.include.view_link,omitempty"`
	ActionEmailInline                         bool    `json:"action.email.inline" url:"action.email.inline"`
	ActionEmailMailserver                     string  `json:"action.email.mailserver,omitempty" url:"action.email.mailserver,omitempty"`
	ActionEmailMaxResults                     int     `json:"action.email.maxresults,omitempty" url:"action.email.maxresults,omitempty"`
	ActionEmailMaxTime                        string  `json:"action.email.maxtime,omitempty" url:"action.email.maxtime,omitempty"`
	ActionEmailMessageAlert                   string  `json:"action.email.message.alert,omitempty" url:"action.email.message.alert,omitempty"`
	ActionEmailMessageReport                  string  `json:"action.email.message.report,omitempty" url:"action.email.message.report,omitempty"`
	ActionEmailPDFView                        string  `json:"action.email.pdfview,omitempty" url:"action.email.pdfview,omitempty"`
	ActionEmailPreprocessResults              string  `json:"action.email.preprocess_results,omitempty" url:"action.email.preprocess_results,omitempty"`
	ActionEmailReportCIDFontList              string  `json:"action.email.reportCIDFontList,omitempty" url:"action.email.reportCIDFontList,omitempty"`
	ActionEmailReportIncludeSplunkLogo        bool    `json:"action.email.reportIncludeSplunkLogo" url:"action.email.reportIncludeSplunkLogo"`
	ActionEmailReportPaperOrientation         string  `json:"action.email.reportPaperOrientation,omitempty" url:"action.email.reportPaperOrientation,omitempty"`
	ActionEmailReportPaperSize                string  `json:"action.email.reportPaperSize,omitempty" url:"action.email.reportPaperSize,omitempty"`
	ActionEmailReportServerEnabled            bool    `json:"action.email.reportServerEnabled" url:"action.email.reportServerEnabled"`
	ActionEmailReportServerURL                string  `json:"action.email.reportServerURL,omitempty" url:"action.email.reportServerURL,omitempty"`
	ActionEmailSendCSV                        int     `json:"action.email.sendcsv,string,omitempty" url:"action.email.sendcsv,omitempty"`
	ActionEmailSendPDF                        bool    `json:"action.email.sendpdf" url:"action.email.sendpdf"`
	ActionEmailSendResults                    bool    `json:"action.email.sendresults" url:"action.email.sendresults"`
	ActionEmailSubject                        string  `json:"action.email.subject,omitempty" url:"action.email.subject,omitempty"`
	ActionEmailTo                             string  `json:"action.email.to,omitempty" url:"action.email.to,omitempty"`
	ActionEmailTrackAlert                     bool    `json:"action.email.track_alert" url:"action.email.track_alert"`
	ActionEmailTTL                            string  `json:"action.email.ttl,omitempty" url:"action.email.ttl,omitempty"`
	ActionEmailUseSSL                         bool    `json:"action.email.use_ssl" url:"action.email.use_ssl"`
	ActionEmailUseTLS                         bool    `json:"action.email.use_tls" url:"action.email.use_tls"`
	ActionEmailWidthSortColumns               bool    `json:"action.email.width_sort_columns" url:"action.email.width_sort_columns"`
	ActionPopulateLookup                      bool    `json:"action.populate_lookup" url:"action.populate_lookup"`
	ActionPopulateLookupCommand               string  `json:"action.populate_lookup.command,omitempty" url:"action.populate_lookup.command,omitempty"`
	ActionPopulateLookupDest                  string  `json:"action.populate_lookup.dest,omitempty" url:"action.populate_lookup.dest,omitempty"`
	ActionPopulateLookupHostname              string  `json:"action.populate_lookup.hostname,omitempty" url:"action.populate_lookup.hostname,omitempty"`
	ActionPopulateLookupMaxResults            int     `json:"action.populate_lookup.maxresults,omitempty" url:"action.populate_lookup.maxresults,omitempty"`
	ActionPopulateLookupMaxTime               int     `json:"action.populate_lookup.maxtime,omitempty" url:"action.populate_lookup.maxtime,omitempty,omitempty"`
	ActionPopulateLookupTrackAlert            bool    `json:"action.populate_lookup.track_alert" url:"action.populate_lookup.track_alert"`
	ActionPopulateLookupTTL                   string  `json:"action.populate_lookup.ttl,omitempty" url:"action.populate_lookup.ttl,omitempty"`
	ActionRSS                                 bool    `json:"-" url:"action.rss"`
	ActionRSSCommand                          string  `json:"action.rss.command,omitempty" url:"action.rss.command,omitempty"`
	ActionRSSHostname                         string  `json:"action.rss.hostname,omitempty" url:"action.rss.hostname,omitempty"`
	ActionRSSMaxResults                       int     `json:"action.rss.maxresults,omitempty" url:"action.rss.maxresults,omitempty"`
	ActionRSSMaxTime                          int     `json:"action.rss.maxtime,omitempty" url:"action.rss.maxtime,omitempty"`
	ActionRSSTrackAlert                       bool    `json:"action.rss.track_alert" url:"action.rss.track_alert"`
	ActionRSSTTL                              string  `json:"action.rss.ttl,omitempty" url:"action.rss.ttl,omitempty"`
	ActionScript                              bool    `json:"-" url:"action.script"`
	ActionScriptCommand                       string  `json:"action.script.command,omitempty" url:"action.script.command,omitempty"`
	ActionScriptFilename                      string  `json:"action.script.filename,omitempty" url:"action.script.filename,omitempty"`
	ActionScriptHostname                      string  `json:"action.script.hostname,omitempty" url:"action.script.hostname,omitempty"`
	ActionScriptMaxResults                    int     `json:"action.script.maxresults,omitempty" url:"action.script.maxresults,omitempty"`
	ActionScriptMaxTime                       int     `json:"action.script.maxtime,omitempty" url:"action.script.maxtime,omitempty"`
	ActionScriptTrackAlert                    bool    `json:"action.script.track_alert" url:"action.script.track_alert"`
	ActionScriptTTL                           string  `json:"action.script.ttl,omitempty" url:"action.script.ttl,omitempty"`
	ActionSummaryIndex                        bool    `json:"action.summary_index,omitempty" url:"action.summary_index"`
	ActionSummaryIndexName                    string  `json:"action.summary_index._name,omitempty" url:"action.summary_index._name,omitempty"`
	ActionSummaryIndexCommand                 string  `json:"action.summary_index.command,omitempty" url:"action.summary_index.command,omitempty"`
	ActionSummaryIndexHostname                string  `json:"action.summary_index.hostname,omitempty" url:"action.summary_index.hostname,omitempty"`
	ActionSummaryIndexInline                  bool    `json:"action.summary_index.inline" url:"action.summary_index.inline"`
	ActionSummaryIndexMaxResults              int     `json:"action.summary_index.maxresults,omitempty" url:"action.summary_index.maxresults,omitempty"`
	ActionSummaryIndexMaxTime                 int     `json:"action.summary_index.maxtime,omitempty" url:"action.summary_index.maxtime,omitempty"`
	ActionSummaryIndexTrackAlert              bool    `json:"action.summary_index.track_alert" url:"action.summary_index.track_alert"`
	ActionSummaryIndexTTL                     string  `json:"action.summary_index.ttl,omitempty" url:"action.summary_index.ttl,omitempty"`
	ActionSlackParamAttachment                string  `json:"action.slack.param.attachment,omitempty" url:"action.slack.param.attachment"`
	ActionSlackParamChannel                   string  `json:"action.slack.param.channel,omitempty" url:"action.slack.param.channel"`
	ActionSlackParamFields                    string  `json:"action.slack.param.fields,omitempty" url:"action.slack.param.fields"`
	ActionSlackParamMessage                   string  `json:"action.slack.param.message,omitempty" url:"action.slack.param.message"`
	ActionSlackParamWebhookUrlOverride        string  `json:"action.slack.param.webhook_url_override,omitempty" url:"action.slack.param.webhook_url_override"`
	ActionJiraServiceDeskParamAccount         string  `json:"action.jira_service_desk.param.account,omitempty" url:"action.jira_service_desk.param.account"`
	ActionJiraServiceDeskParamJiraProject     string  `json:"action.jira_service_desk.param.jira_project,omitempty" url:"action.jira_service_desk.param.jira_project"`
	ActionJiraServiceDeskParamJiraIssueType   string  `json:"action.jira_service_desk.param.jira_issue_type,omitempty" url:"action.jira_service_desk.param.jira_issue_type"`
	ActionJiraServiceDeskParamJiraSummary     string  `json:"action.jira_service_desk.param.jira_summary,omitempty" url:"action.jira_service_desk.param.jira_summary"`
	ActionJiraServiceDeskParamJiraPriority    string  `json:"action.jira_service_desk.param.jira_priority,omitempty" url:"action.jira_service_desk.param.jira_priority"`
	ActionJiraServiceDeskParamJiraDescription string  `json:"action.jira_service_desk.param.jira_description,omitempty" url:"action.jira_service_desk.param.jira_description"`
	ActionWebhookParamUrl                     string  `json:"action.webhook.param.url,omitempty" url:"action.webhook.param.url"`
	AlertDigestMode                           bool    `json:"alert.digest_mode" url:"alert.digest_mode"`
	AlertExpires                              string  `json:"alert.expires,omitempty" url:"alert.expires,omitempty"`
	AlertSeverity                             int     `json:"alert.severity,omitempty" url:"alert.severity,omitempty"`
	AlertSuppress                             bool    `json:"alert.suppress" url:"alert.suppress"`
	AlertSuppressFields                       string  `json:"alert.suppress.fields,omitempty" url:"alert.suppress.fields,omitempty"`
	AlertSuppressPeriod                       string  `json:"alert.suppress.period,omitempty" url:"alert.suppress.period,omitempty"`
	AlertTrack                                bool    `json:"alert.track" url:"alert.track"`
	AlertComparator                           string  `json:"alert_comparator,omitempty" url:"alert_comparator,omitempty"`
	AlertCondition                            string  `json:"alert_condition,omitempty" url:"alert_condition,omitempty"`
	AlertThreshold                            string  `json:"alert_threshold,omitempty" url:"alert_threshold,omitempty"`
	AlertType                                 string  `json:"alert_type,omitempty" url:"alert_type,omitempty"`
	AllowSkew                                 string  `json:"allow_skew,omitempty" url:"allow_skew,omitempty"`
	AutoSummarize                             bool    `json:"auto_summarize,omitempty" url:"auto_summarize,omitempty"`
	AutoSummarizeCommand                      string  `json:"auto_summarize.command,omitempty" url:"auto_summarize.command,omitempty"`
	AutoSummarizeCronSchedule                 string  `json:"auto_summarize.cron_schedule,omitempty" url:"auto_summarize.cron_schedule,omitempty"`
	AutoSummarizeDispatchEarliestTime         string  `json:"auto_summarize.dispatch.earliest_time,omitempty" url:"auto_summarize.dispatch.earliest_time,omitempty"`
	AutoSummarizeDispatchLatestTime           string  `json:"auto_summarize.dispatch.latest_time,omitempty" url:"auto_summarize.dispatch.latest_time,omitempty"`
	AutoSummarizeDispatchTimeFormat           string  `json:"auto_summarize.dispatch.time_format,omitempty" url:"auto_summarize.dispatch.time_format,omitempty"`
	AutoSummarizeDispatchTTL                  string  `json:"auto_summarize.dispatch.ttl,omitempty" url:"auto_summarize.dispatch.ttl,omitempty"`
	AutoSummarizeMaxDisabledBuckets           int     `json:"auto_summarize.max_disabled_buckets,omitempty" url:"auto_summarize.max_disabled_buckets,omitempty"`
	AutoSummarizeMaxSummaryRatio              float64 `json:"auto_summarize.max_summary_ratio,omitempty" url:"auto_summarize.max_summary_ratio,omitempty"`
	AutoSummarizeMaxSummarySize               int     `json:"auto_summarize.max_summary_size,omitempty" url:"auto_summarize.max_summary_size,omitempty"`
	AutoSummarizeMaxTime                      int     `json:"auto_summarize.max_time,omitempty" url:"auto_summarize.max_time,omitempty"`
	AutoSummarizeSuspendPeriod                string  `json:"auto_summarize.suspend_period,omitempty" url:"auto_summarize.suspend_period,omitempty"`
	AutoSummarizeTimespan                     string  `json:"auto_summarize.timespan,omitempty" url:"auto_summarize.timespan,omitempty"`
	CronSchedule                              string  `json:"cron_schedule,omitempty" url:"cron_schedule,omitempty"`
	Description                               string  `json:"description,omitempty" url:"description,omitempty"`
	Disabled                                  bool    `json:"disabled" url:"disabled"`
	DispatchBuckets                           int     `json:"dispatch.buckets,omitempty" url:"dispatch.buckets,omitempty"`
	DispatchEarliestTime                      string  `json:"dispatch.earliest_time,omitempty" url:"dispatch.earliest_time,omitempty"`
	DispatchIndexEarliest                     string  `json:"dispatch.index_earliest,omitempty" url:"dispatch.index_earliest,omitempty"`
	DispatchIndexLatest                       string  `json:"dispatch.index_latest,omitempty" url:"dispatch.index_latest,omitempty"`
	DispatchIndexedRealtime                   bool    `json:"dispatch.indexedRealtime" url:"dispatch.indexedRealtime"`
	DispatchIndexedRealtimeOffset             int     `json:"dispatch.indexedRealtimeOffset" url:"dispatch.indexedRealtimeOffset,omitempty"`
	DispatchIndexedRealtimeMinspan            int     `json:"dispatch.indexedRealtimeMinspan" url:"dispatch.indexedRealtimeMinspan,omitempty"`
	DispatchLatestTime                        string  `json:"dispatch.latest_time,omitempty" url:"dispatch.latest_time,omitempty"`
	DispatchLookups                           bool    `json:"dispatch.lookups" url:"dispatch.lookups"`
	DispatchMaxCount                          int     `json:"dispatch.max_count,omitempty" url:"dispatch.max_count,omitempty"`
	DispatchMaxTime                           int     `json:"dispatch.max_time,omitempty" url:"dispatch.max_time,omitempty"`
	DispatchReduceFreq                        int     `json:"dispatch.reduce_freq,omitempty" url:"dispatch.reduce_freq,omitempty"`
	DispatchRtBackfill                        bool    `json:"dispatch.rt_backfill" url:"dispatch.rt_backfill"`
	DispatchRtMaximumSpan                     int     `json:"dispatch.rt_maxtimespan" url:"dispatch.rt_maxtimespan,omitempty"`
	DispatchSpawnProcess                      bool    `json:"dispatch.spawn_process" url:"dispatch.spawn_process"`
	DispatchTimeFormat                        string  `json:"dispatch.time_format,omitempty" url:"dispatch.time_format,omitempty"`
	DispatchTTL                               string  `json:"dispatch.ttl,omitempty" url:"dispatch.ttl,omitempty"`
	DisplayView                               string  `json:"displayview,omitempty" url:"displayview,omitempty"`
	IsScheduled                               bool    `json:"is_scheduled" url:"is_scheduled"`
	IsVisible                                 bool    `json:"is_visible" url:"is_visible"`
	MaxConcurrent                             int     `json:"max_concurrent,omitempty" url:"max_concurrent,omitempty"`
	NextScheduledTime                         string  `json:"next_scheduled_time,omitempty" url:"next_scheduled_time,omitempty"`
	QualifiedSearch                           string  `json:"qualifiedSearch,omitempty" url:"qualifiedSearch,omitempty"`
	RealtimeSchedule                          bool    `json:"realtime_schedule" url:"realtime_schedule"`
	RequestUIDispatchApp                      string  `json:"request.ui_dispatch_app,omitempty" url:"request.ui_dispatch_app,omitempty"`
	RequestUIDispatchView                     string  `json:"request.ui_dispatch_view,omitempty" url:"request.ui_dispatch_view,omitempty"`
	RestartOnSearchPeerAdd                    bool    `json:"restart_on_searchpeer_add" url:"restart_on_searchpeer_add"`
	RunOnStartup                              bool    `json:"run_on_startup" url:"run_on_startup"`
	ScheduleWindow                            string  `json:"schedule_window,omitempty" url:"schedule_window,omitempty"`
	SchedulePriority                          string  `json:"schedule_priority,omitempty" url:"schedule_priority,omitempty"`
	Search                                    string  `json:"search,omitempty" url:"search,omitempty"`
	VSID                                      string  `json:"vsid,omitempty" url:"vsid,omitempty"`
	WorkloadPool                              string  `json:"workload_pool,omitempty" url:"workload_pool,omitempty"`
}

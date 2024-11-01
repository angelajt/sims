// Code generated by "core generate -add-types"; DO NOT EDIT.

package main

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "main.Config", IDName: "config", Doc: "Config has config parameters related to running the sim", Fields: []types.Field{{Name: "NRuns", Doc: "total number of runs to do when running Train"}, {Name: "NEpochs", Doc: "total number of epochs per run"}, {Name: "NZero", Doc: "stop run after this number of perfect, zero-error epochs."}, {Name: "TestInterval", Doc: "how often to run through all the test patterns, in terms of training epochs.\ncan use 0 or -1 for no testing."}}})

var _ = types.AddType(&types.Type{Name: "main.Sim", IDName: "sim", Doc: "Sim encapsulates the entire simulation model, and we define all the\nfunctionality as methods on this struct.  This structure keeps all relevant\nstate information organized and available without having to pass everything around\nas arguments to methods, and provides the core GUI interface (note the view tags\nfor the fields which provide hints to how things should be displayed).", Fields: []types.Field{{Name: "Config", Doc: "Config contains misc configuration parameters for running the sim"}, {Name: "Net", Doc: "the network -- click to view / edit parameters for layers, paths, etc"}, {Name: "Params", Doc: "network parameter management"}, {Name: "Loops", Doc: "contains looper control loops for running sim"}, {Name: "Stats", Doc: "contains computed statistic values"}, {Name: "Logs", Doc: "Contains all the logs and information about the logs.'"}, {Name: "Envs", Doc: "Environments"}, {Name: "Context", Doc: "leabra timing parameters and state"}, {Name: "ViewUpdate", Doc: "netview update parameters"}, {Name: "GUI", Doc: "manages all the gui elements"}, {Name: "RandSeeds", Doc: "a list of random seeds to use for each run"}}})

var _ = types.AddType(&types.Type{Name: "main.SentGenEnv", IDName: "sent-gen-env", Doc: "SentGenEnv generates sentences using a grammar that is parsed from a\ntext file.  The core of the grammar is rules with various items\nchosen at random during generation -- these items can be\nmore rules terminal tokens.", Fields: []types.Field{{Name: "Name", Doc: "name of this environment"}, {Name: "Rules", Doc: "core sent-gen rules -- loaded from a grammar / rules file -- Gen() here generates one sentence"}, {Name: "PPassive", Doc: "probability of generating passive sentence forms"}, {Name: "WordTrans", Doc: "translate unambiguous words into ambiguous words"}, {Name: "Words", Doc: "list of words used for activating state units according to index"}, {Name: "WordMap", Doc: "map of words onto index in Words list"}, {Name: "Roles", Doc: "list of roles used for activating state units according to index"}, {Name: "RoleMap", Doc: "map of roles onto index in Roles list"}, {Name: "Fillers", Doc: "list of filler concepts used for activating state units according to index"}, {Name: "FillerMap", Doc: "map of roles onto index in Words list"}, {Name: "AmbigVerbs", Doc: "ambiguous verbs"}, {Name: "AmbigNouns", Doc: "ambiguous nouns"}, {Name: "AmbigVerbsMap", Doc: "map of ambiguous verbs"}, {Name: "AmbigNounsMap", Doc: "map of ambiguous nouns"}, {Name: "CurSentOrig", Doc: "original current sentence as generated from Rules"}, {Name: "CurSent", Doc: "current sentence, potentially transformed to passive form"}, {Name: "NAmbigNouns", Doc: "number of ambiguous nouns"}, {Name: "NAmbigVerbs", Doc: "number of ambiguous verbs (0 or 1)"}, {Name: "SentInputs", Doc: "generated sequence of sentence inputs including role-filler queries"}, {Name: "SentIndex", Doc: "current index within sentence inputs"}, {Name: "QType", Doc: "current question type -- from 4th value of SentInputs"}, {Name: "WordState", Doc: "current sentence activation state"}, {Name: "RoleState", Doc: "current role query activation state"}, {Name: "FillerState", Doc: "current filler query activation state"}, {Name: "Seq", Doc: "sequence counter within epoch"}, {Name: "Tick", Doc: "tick counter within sequence"}, {Name: "Trial", Doc: "trial is the step counter within sequence - how many steps taken within current sequence -- it resets to 0 at start of each sequence"}}})

var _ = types.AddType(&types.Type{Name: "main.ProbeEnv", IDName: "probe-env", Doc: "ProbeEnv generates sentences using a grammar that is parsed from a\ntext file.  The core of the grammar is rules with various items\nchosen at random during generation -- these items can be\nmore rules terminal tokens.", Fields: []types.Field{{Name: "Name", Doc: "name of this environment"}, {Name: "Words", Doc: "list of words used for activating state units according to index"}, {Name: "WordState", Doc: "current sentence activation state"}, {Name: "Trial", Doc: "trial is the step counter within sequence - how many steps taken\nwithin current sequence -- it resets to 0 at start of each sequence."}}})
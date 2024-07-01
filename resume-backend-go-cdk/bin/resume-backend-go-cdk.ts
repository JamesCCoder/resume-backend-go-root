#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { ResumeBackendGoCdkStack } from "../lib/resume-backend-go-cdk-stack";

const app = new cdk.App();
new ResumeBackendGoCdkStack(app, 'ResumeBackendGoCdkStack');

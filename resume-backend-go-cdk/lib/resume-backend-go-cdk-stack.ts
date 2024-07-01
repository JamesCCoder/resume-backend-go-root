import * as cdk from '@aws-cdk/core';
import * as lambda from '@aws-cdk/aws-lambda';
import * as apigateway from '@aws-cdk/aws-apigateway';
import * as sfn from '@aws-cdk/aws-stepfunctions';
import * as tasks from '@aws-cdk/aws-stepfunctions-tasks';
import * as path from 'path';

export class ResumeBackendGoCdkStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // 定义 Lambda 函数
    const backendLambda = new lambda.Function(this, 'BackendLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2,
      handler: 'main',
      code: lambda.Code.fromAsset(path.join(__dirname, '..', '..', 'resume-backend-go', 'build')),
      environment: {
        // 在此设置环境变量，例如 MongoDB URI
        MONGODB_URI: process.env.MONGODB_URI || "mongodb+srv://admin:19890417Qq@cluster0.anw7czf.mongodb.net/?appName=Cluster0"
      }
    });

    // 定义 API Gateway 并集成 Lambda 函数
    const api = new apigateway.LambdaRestApi(this, 'ResumeBackendApi', {
      handler: backendLambda,
      proxy: false
    });

    const items = api.root.addResource('items');
    items.addMethod('GET');  // GET /items
    items.addMethod('POST'); // POST /items

    const singleItem = items.addResource('{id}');
    singleItem.addMethod('GET');  // GET /items/{id}
    singleItem.addMethod('PUT');  // PUT /items/{id}
    singleItem.addMethod('DELETE'); // DELETE /items/{id}

    // 定义 Step Function 任务
    const stepFunctionTask = new tasks.LambdaInvoke(this, 'InvokeBackendLambda', {
      lambdaFunction: backendLambda,
      outputPath: '$.Payload',
    });

    // 定义 Step Function 状态机
    const definition = stepFunctionTask;
    const stateMachine = new sfn.StateMachine(this, 'StateMachine', {
      definition,
      timeout: cdk.Duration.minutes(5),
    });

    // 输出 API Gateway 的 URL
    new cdk.CfnOutput(this, 'ApiUrl', {
      value: api.url,
    });

    // 输出 Step Function 的 ARN
    new cdk.CfnOutput(this, 'StateMachineArn', {
      value: stateMachine.stateMachineArn,
    });
  }
}

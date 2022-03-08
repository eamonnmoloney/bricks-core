FROM public.ecr.aws/lambda/go

WORKDIR ${LAMBDA_TASK_ROOT}

COPY bricks-core ${LAMBDA_TASK_ROOT}

CMD [ "bricks-core" ]
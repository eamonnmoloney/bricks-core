FROM public.ecr.aws/lambda/go

ENV MODE production

WORKDIR ${LAMBDA_TASK_ROOT}

COPY bricks-core ${LAMBDA_TASK_ROOT}

CMD [ "bricks-core" ]